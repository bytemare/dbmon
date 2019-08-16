package main

import (
	"context"
	"github.com/bytemare/dbmon/dbmon"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

const addr = "localhost"
const port = "4000"
const clusterId = "roachy"

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := dbmon.NewHealCheckClient(conn)

	// Contact the server and print out its response.

	time.Sleep(7 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Pull(ctx, &dbmon.PullRequest{
		ClusterId:            clusterId,
		NbProbes:             1,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	if err != nil {
		log.Fatalf("received error : %s", err)
	}
	//log.Printf("Client received: %s", r.Reports)

	probes := r.Probes

	log.Infof("Client received a response !")
	log.Infof("problen : %d", len(probes))
	for i, p := range probes {
		log.Infof("probe %d : status %s", i, p.Status)
		log.Infof("time  %s", p.Timestamp)
		log.Infof("response len %d", p.Len)
		log.Infof("payload (%d/%d): '%s'", len(p.Data[i]), len(string(p.Data[i])), string(p.Data[i]))
	}
	log.Infof("Client finished.")
}
