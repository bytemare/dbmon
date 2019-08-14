package main

import (
	"github.com/bytemare/dbmon"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

const port  = ":8081"

func main() {

	mon := dbmon.NewDBMon(port)

	mon.Start()
	log.Info("Server started.")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	mon.Stop()

	log.Info("Server stopped.")
}