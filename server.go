package dbmon

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/bytemare/dbmon/dbmon"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"sync"
)

// DBMon
type DBMon struct {
	// TODO this may not be very memrory efficient, look for a better solution
	clusters map[string]*Cluster
	data     map[string][]*pb.Report // Cache to hold data to serve per cluster : a list of reports
	mux      sync.Mutex              // Mutex to protect against concurrent read/write on cache
	source   <-chan *pb.Report       // Channel through which data is arriving
	sync     chan struct{}           // Stop signal channel
	port     string                  // Network port to listen on
	grpc     *grpc.Server            // The gRPC Server Handle
}

// NewDBMon initialises and returns a new DBMon struct
func NewDBMon(port string, serverChan <-chan *pb.Report) *DBMon {
	return &DBMon{
		clusters: make(map[string]*Cluster),
		data:     make(map[string][]*pb.Report),
		mux:      sync.Mutex{},
		source:   serverChan,
		sync:     make(chan struct{}),
		port:     port,
		grpc:     grpc.NewServer(),
	}
}

// Start starts the DBMon server
func (mon *DBMon) Start() {

	lis, err := net.Listen("tcp", mon.port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterHealCheckServer(mon.grpc, mon)
	if err := mon.grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

loop:
	for {
		select {

		case <-mon.sync:
			log.Info("Server shutting down.")
			break loop

		case r := <-mon.source:
			log.Info("Server received a new report")
			mon.mux.Lock()
			mon.data[r.ClusterID] = append(mon.data[r.ClusterID], r)
		}
	}

}

// RegisterCluster adds a new cluster to serve reports for
func (mon *DBMon) RegisterCluster(cluster *Cluster) error {
	if _, ok := mon.clusters[cluster.id]; ok {
		log.Errorf("Cluster '%s' has already been registered.", cluster.id)
		return errors.New(fmt.Sprintf("Cluster registration failed : value already exists '%s'.", cluster.id))
	} else {
		mon.clusters[cluster.id] = cluster
		mon.data[cluster.id] = []*pb.Report{}
		return nil
	}
}

// Returns the list of strings representing reports
func (mon *DBMon) extractReports(clusterID string) []*pb.Report {

	mon.mux.Lock()
	reports := mon.data[clusterID]
	output := make([]*pb.Report, len(reports))
	copy(output, reports)
	delete(mon.data, clusterID)
	mon.mux.Unlock()

	return output
}

// Stop stops the DBMon server
func (mon *DBMon) Stop() {
	mon.sync <- struct{}{}
	mon.grpc.GracefulStop()
}

/**
** Implement the gRPC
 */

// SayHello implements HealthCheckerServer
func (mon DBMon) Pull(ctx context.Context, in *pb.PullRequest) (*pb.PullReply, error) {
	log.Info("Received pull for : %v", in.ClusterId)

	reply := &pb.PullReply{
		ClusterId:            in.ClusterId,
		Reports:              nil,
		Error:                nil,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	// Verify if cluster is registered
	if _, ok := mon.clusters[in.ClusterId]; !ok {
		/*err := &pb.Error{
			Code:                 uint32(5),
			Message:              "This cluster is not registered",
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
		reply.Error = err*/
		err1 := status.Error(codes.NotFound, "id was not found")
		return nil, err1
		//return reply, errors.New(err.Message)
	}

	reply.Reports = mon.extractReports(in.ClusterId)

	return reply, nil
}
