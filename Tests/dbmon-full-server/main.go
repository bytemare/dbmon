package main

import (
	"github.com/bytemare/dbmon"
	"github.com/bytemare/dbmon/connectors"
	dbmon2 "github.com/bytemare/dbmon/dbmon"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

const addr = "http://localhost"
const serverPort = ":4000"
const clusterPort  = "8080"
const cludterId = "roachy"

func main() {

	source := make(chan *dbmon2.Report)

	// Set up server
	mon := dbmon.NewDBMon(serverPort, source)
	go mon.Start()
	log.Info("Server started.")

	// Set up collector
	roachCon := connectors.NewConnector(addr, clusterPort)
	cluster := dbmon.NewCluster(cludterId, roachCon)
	collector := dbmon.NewCollector(source)


	log.Info("Collection set.")

	// It is important to first add a cluster in the server
	_ = mon.RegisterCluster(cluster)


	log.Info("Cluster registered to server.")

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
