package main

import (
	"context"
	"github.com/bytemare/dbmon/dbmon"
	"google.golang.org/grpc"
	"log"
	"time"
)

const addr = "localhost"
const port = "4000"
const clusterId = "roachy"

func main()  {

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr + ":" + port, grpc.WithInsecure())
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
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	if err != nil {
		log.Fatalf("received error : %s", err)
	}
	log.Printf("Client received: %s", r.Reports)

}