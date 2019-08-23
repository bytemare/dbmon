package dbmon

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"sync"
)

// DBMon holds the cache of reports send from the collector to serve them to clients
type DBMon struct {
	// TODO this may not be very memory efficient, look for a better solution
	clusters map[string]*Cluster
	data     map[string][]*Probe // Cache to hold data to serve per cluster : a list of reports
	mux      sync.Mutex          // Mutex to protect against concurrent read/write on cache
	source   <-chan *Probe       // Channel through which data is arriving
	sync     chan struct{}       // Stop signal channel
	port     string              // Network port to listen on
	grpc     *grpc.Server        // The gRPC Server Handle
}

// NewDBMon initialises and returns a new DBMon struct
func NewDBMon(port string, serverChan <-chan *Probe) *DBMon {
	return &DBMon{
		clusters: make(map[string]*Cluster),
		data:     make(map[string][]*Probe),
		mux:      sync.Mutex{},
		source:   serverChan,
		sync:     make(chan struct{}),
		port:     port,
		grpc:     grpc.NewServer(),
	}
}

// Start starts the DBMon server
func (mon *DBMon) Start() {

	log.Info("Starting dbmon.")

	lis, err := net.Listen("tcp", mon.port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Info("dbmon listening.")

	RegisterHealCheckServer(mon.grpc, mon)

	log.Info("dbmon registered to grpc.")

	/* Goroutine to fetch incoming probes */
	go func() {
	loop:
		for {
			select {

			case <-mon.sync:
				log.Info("Server shutting down.")
				break loop

			case r := <-mon.source:
				log.Info("Server received a new probe for ", r.ClusterID)
				mon.mux.Lock()
				// TODO : this here will work fine as long as the collector runs with the server
				// TODO : but if they are split, we have to ensure that the clusterid is still registered in the server
				probes := mon.data[r.ClusterID]
				probes = append(probes, r)
				mon.data[r.ClusterID] = probes
				mon.mux.Unlock()
			}
		}
	}()

	log.Info("Server routine launched. Starting server.")

	if err := mon.grpc.Serve(lis); err != nil {
		mon.Stop()
		log.Errorf("failed to serve: %v", err)
	}
}

// RegisterCluster adds a new cluster to serve reports for
func (mon *DBMon) RegisterCluster(cluster *Cluster) error {
	if _, ok := mon.clusters[cluster.id]; ok {
		log.Errorf("Cluster '%s' has already been registered.", cluster.id)
		return fmt.Errorf("cluster registration failed : value already exists '%s'", cluster.id)
	}

	log.Infof("Registering new cluster '%s' to server", cluster.id)

	mon.mux.Lock()
	mon.clusters[cluster.id] = cluster
	mon.data[cluster.id] = []*Probe{}
	mon.mux.Unlock()
	return nil
}

// Returns the list of strings representing reports
func (mon *DBMon) extractProbes(clusterID string, nbProbes int) []*Probe {

	mon.mux.Lock()
	cache := mon.data[clusterID]

	if len(cache) == 0 {
		mon.mux.Unlock()
		return nil
	}

	if nbProbes <= 0 {
		nbProbes = len(cache)
	}
	probes := make([]*Probe, nbProbes)
	copy(probes[:nbProbes], cache[:nbProbes])

	// Delete head elements : this method avoids memory leaks
	copy(cache, cache[nbProbes:])
	for i := nbProbes; i < len(cache); i++ {
		cache[i] = nil
	}
	cache = cache[:len(cache)-nbProbes]

	mon.mux.Unlock()

	return probes
}

// Stop stops the DBMon server
func (mon *DBMon) Stop() {
	mon.sync <- struct{}{}
	mon.grpc.GracefulStop()
}

/**
** Implement the gRPC interface
 */

// Pull implements HealthCheckerServer
func (mon *DBMon) Pull(ctx context.Context, in *PullRequest) (*PullReply, error) {
	log.Infof("Received pull for : %s", in.ClusterId)

	// Verify if cluster is registered
	if _, ok := mon.clusters[in.ClusterId]; !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Cluster was not found : '%s'", in.ClusterId))
	}

	reply := &PullReply{
		ClusterId:            in.ClusterId,
		Probes:               mon.extractProbes(in.ClusterId, int(in.NbProbes)),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	return reply, nil
}
