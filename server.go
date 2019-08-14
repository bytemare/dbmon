// dbmon is a server daemon with a gRPC API endpoint that serves information about a database cluster
package dbmon

import (
	"context"
	pb "github.com/bytemare/dbmon/dbmon"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

// DBMon
type DBMon struct {
	data   map[*Cluster]report // Cache to hold data to serve
	source <-chan [2]string    // Channel through which data is arriving
	port	string				// Network port to listen on
	grpc	*grpc.Server		// The gRPC Server Handle
}

// report holds a cluster's accumulated information
type report struct {
	clusterID *Cluster
	data      []string
}

// NewDBMon initialises and returns a new DBMon struct
func NewDBMon(port string) *DBMon {
	return &DBMon{
		data:   make(map[*Cluster]report),
		source: make(chan [2]string),
		port:port,
		grpc: grpc.NewServer(),
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
}

// Stop stops the DBMon server
func (mon *DBMon) Stop() {
	mon.grpc.GracefulStop()
}

/**
** Implement the gRPC
 */

// SayHello implements HealthCheckerServer
func (mon DBMon) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}