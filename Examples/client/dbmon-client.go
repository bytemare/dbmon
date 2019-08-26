package main

/**
TODO : apparently, the grpc client doesn't implement a call/request based timeout, but only a total connection limit.
	So we would have to hold a counter, resetting it to 0 after each successful call, to trace if the timeout is due to
	a call or the whole connection
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bytemare/dbmon"
	cockroachdb "github.com/bytemare/dbmon/connectors/CockroachDB"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const addr = "localhost"
const port = "4000"
const clusterID = "roachy"
const nbProbes = 0
const refresh = 2 * time.Second
const timeout = 20 * time.Second

/*
func printNode(node *cockroachdb.NodeStats) {
	log.Infof("Node Info : %d", node.ID)
	log.Infof("Ranges: %d", node.Ranges)
	log.Infof("Replicas : %d", node.Replicas)
	log.Infof("Cap. Usage : %d", node.CapacityUsage)
	log.Infof("Cap. Reserved : %d", node.CapacityReserved)
	log.Infof("Cap. Available : %d", node.CapacityAvailable)
	log.Infof("Cap. : %d", node.Capacity)
	log.Infof("SystemMem : %d", node.TotalSystemMemory)
	log.Infof("Q/s %f", node.QueriesPerSecond)
	log.Infof("Heartbeat Latency : %d", node.Latency99)
	log.Infof("ClockOffset : %d", node.ClockOffset)
}
*/

func setupConnection() *connection {

	connection := &connection{
		Conn:   nil,
		Client: nil,
		Ctx:    nil,
		Cancel: nil,
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	connection.Conn = conn
	connection.Client = dbmon.NewHealCheckClient(connection.Conn)
	clientDeadline := time.Now().Add(timeout)
	//ctx, cancel := context.WithTimeout(context.Background(), timeout)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	connection.Ctx = ctx
	connection.Cancel = cancel

	return connection
}

type connection struct {
	Conn   *grpc.ClientConn
	Client dbmon.HealCheckClient
	Ctx    context.Context
	Cancel context.CancelFunc
}

func (c *connection) close() {
	_ = c.Conn.Close()
	c.Cancel()
}

// Pull implements the grpc client interface
func (c *connection) Pull() (*dbmon.PullReply, error) {

	reply, err := c.Client.Pull(c.Ctx, &dbmon.PullRequest{
		ClusterId:            clusterID,
		NbProbes:             nbProbes,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})

	if err != nil {
		log.Errorf("received error : %s", err)
		return nil, err
	}

	return reply, nil
}

func integerByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func analyse(clusterID string, timestamp string, nodes []*cockroachdb.NodeStats) *analysis {
	a := &analysis{
		clusterID:           clusterID,
		timestamp:           timestamp,
		nodes:               len(nodes),
		usedCap:             "",
		availableCap:        "",
		queries:             0,
		heartbeatP99:        "",
		ranges:              0,
		highestOffset:       0,
		highestOffsetNodeID: 0,
	}

	var usedCap uint64
	var availCap int64

	for _, node := range nodes {
		usedCap += node.CapacityUsage
		availCap += int64(node.CapacityAvailable)
		a.queries += node.QueriesPerSecond
		if node.Ranges > a.ranges {
			a.ranges = node.Ranges
		}
		if node.ClockOffset > a.highestOffset {
			a.highestOffset = node.ClockOffset
			a.highestOffsetNodeID = node.ID
		}
	}

	// TODO P99

	a.usedCap = fmt.Sprintf("%.2f%%", 100*float64(usedCap)/float64(availCap))
	a.availableCap = integerByteCountSI(availCap)

	return a
}

func pullPrint(c *connection) error {
	// Contact the server
	reply, err := c.Pull()
	if err != nil {
		return err
	}

	//log.Infof("Client received a response !")

	probes := reply.Probes
	var nodes []*cockroachdb.NodeStats

	for _, p := range probes {
		err = json.Unmarshal(p.Data, &nodes)
		if err != nil {
			log.Error("Error in json : ", err)
			continue
		}

		displayConsole(analyse(p.ClusterID, time.Now().Format("2019-08-16 23:43:04:4675"), nodes))
	}

	return nil
}

func main() {

	log.Warnf("This is a demo with default values, and the deadline is set to %s", timeout.String())

	// Set up a connection
	connection := setupConnection()
	defer connection.close()

	// Sync
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	stop := make(chan struct{})

	//time.Sleep(3 * time.Second)

	ticker := time.NewTicker(refresh)

	go func() {

		for {
			select {

			case <-stop:
				return

			case <-ticker.C:
				err := pullPrint(connection)
				if err != nil {
					sigs <- syscall.SIGINT
				}
			}
		}
	}()

	// Intercept interruption

	log.Info("Waiting for signal.")
	s := <-sigs
	log.Info("Signal received : ", s.String())

	stop <- struct{}{}
	close(stop)
	ticker.Stop()

	log.Infof("Client finished.")
}

/**
*
*
*
*
*
 */

const (
	clearConsole = "\x1Bc"
	lineStart    = "\t\t> "
	topLine      = green + "[dbmon]" + blue + " \tCluster :  %s" + stop + "\t\tLast updated : %s"
	//noData			= tabSpace + "\t--- No available data ---"
	summary        = "Health check summary :"
	summaryNodes   = "Nodes : %d"
	summaryUsedCap = "Used Capacity : %s / %s"
	summaryQueries = "Queries per second : %.2f"
	summaryP99     = "Heartbeat P99 Latency : %s ms"
	database       = "Ranges : %d"
	highestOffset  = "Highest Clock offset %d ns ( node %d )"

	// ANSI Colours
	//red   = "\033[31;1;1m"
	green = "\033[32m"
	blue  = "\033[34m"
	stop  = "\033[0m"
)

type analysis struct {
	clusterID           string
	timestamp           string // time.Now().Format("2019-08-17 23-03-18:745")
	nodes               int
	usedCap             string
	availableCap        string
	queries             float64
	heartbeatP99        string
	ranges              uint
	highestOffset       uint32
	highestOffsetNodeID int
}

func displayConsole(a *analysis) {
	var output string

	output += fmt.Sprintf(topLine+"\n", a.clusterID, a.timestamp)
	output += fmt.Sprint(summary) + "\n"
	output += lineStart + fmt.Sprintf(summaryNodes, a.nodes) + "\n"
	output += lineStart + fmt.Sprintf(summaryUsedCap, a.usedCap, a.availableCap) + "\n"
	output += lineStart + fmt.Sprintf(summaryQueries, a.queries) + "\n"
	output += lineStart + fmt.Sprintf(summaryP99, a.heartbeatP99) + "\n"
	output += lineStart + fmt.Sprintf(database, a.ranges) + "\n"
	output += lineStart + fmt.Sprintf(highestOffset, a.highestOffset, a.highestOffsetNodeID) + "\n"

	fmt.Print(clearConsole)
	fmt.Print(output)
}
