package main

import (
	"github.com/bytemare/dbmon"
	"github.com/bytemare/dbmon/connectors/CockroachDB"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

const addr = "http://localhost"
const serverPort = ":4000"
const clusterPort = "8080"
const clusterId = "roachy"

func main() {

	source := make(chan *dbmon.Probe)

	// Set up components
	mon := dbmon.NewDBMon(serverPort, source)
	collector := dbmon.NewCollector(source)
	roachCon := cockroachdb.NewConnector(clusterId, addr, clusterPort)
	cluster := dbmon.NewCluster(clusterId, roachCon)

	// Register Cluster to server
	// todo : Clusters have to be added to server before it starts, and before running the collector,
	//  but can only be added to collector after it's running...
	_ = mon.RegisterCluster(cluster)
	log.Info("Cluster registered to server.")

	// Run server
	go mon.Start()

	// It is less problematic to add a cluster to the collector if it's already running
	go collector.Start()

	log.Info("Collector started.")
	collector.RegisterNewCluster(cluster)
	log.Info("Cluster registered to collector.")

	// Intercept interruption
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Waiting for signal.")
	s := <-sigs

	log.Info("Signal received : ", s.String())

	// Shut everything down
	collector.Stop()
	mon.Stop()

	log.Info("Everything shut down. Good night.")
}
