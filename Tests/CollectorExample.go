package main

import (
	"../dbmon"
	"github.com/bytemare/dbmon/Tests/CockroachDB"
	log "github.com/sirupsen/logrus"
)



// An example program using the collector
func main() {

	roachCon := CockroachDB.NewConnector("http://localhost", "8080")

	cluster := dbmon.NewCluster("roachy", *roachCon)

	serverChan := make(chan [2]string)

	collector := dbmon.NewCollector(serverChan)

	go collector.Start()

	collector.RegisterNewCluster(cluster)

	for i := 0 ; i < 5 ; i++  {
		log.Info("Received from collector ", <-serverChan)
	}

	collector.Stop()

	close(serverChan)

}
