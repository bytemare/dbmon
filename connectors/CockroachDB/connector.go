// Package cockroachdb implements the dbmon Connector interface
// It hols everything needed to connect to a CockroachDB cluster and retrieve data
package cockroachdb

/**
TODO : The json API is subject to change, so we're sending the raw json response, for now.
TODO : Maybe using the cockroachdb command could help : but how would we do this with a remote cluster ?
*/

import (
	"encoding/json"
	"errors"
	"github.com/bytemare/dbmon"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	roachPeriod  = 1 * time.Second
	roachTimeout = 10 * time.Second
)

const (
	reqAllNodes = "/_status/nodes"
	/*
		reqHealth         = "/health"
		reqHealthReady    = "/health?ready=1"
		reqSingleNode     = "/_status/nodes/[node_id]"
		reqSingleHotrange = "_status/hotranges?node_id=[node_id]"
		reqAllHotRanges   = "/_status/hotranges"
		reqClusterRaft    = "/_status/raft"
		reqClusterRange   = "/_status/range/1"
	*/
)

// RoachConnector is the struct to interact with, holding connection info about the target cluster
type RoachConnector struct {
	clusterID string
	ip        string
	port      string
	period    time.Duration // Period to repetitively interrogate the cluster
	timeout   time.Duration // Timeout per request
	netClient *http.Client
}

// NewConnector returns a new initialised connector for a CockRoachDB endpoint
func NewConnector(id string, ip string, port string, period time.Duration, timeout time.Duration) *RoachConnector {

	if period == 0 {
		period = roachPeriod
	}
	if timeout == 0 {
		timeout = roachTimeout
	}
	return &RoachConnector{
		clusterID: id,
		ip:        ip,
		port:      port,
		period:    period,
		timeout:   timeout,
		netClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// get operates a HTTP get request to the given endpoint, and returns the Response on success (else nil)
func (rc RoachConnector) get(endpoint string) (*http.Response, error) {
	resp, err := rc.netClient.Get(endpoint)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return resp, nil
}

func extractNodeInfo(nodes *AllNodes) []*NodeStats {
	nodeStats := make([]*NodeStats, 0, len(nodes.Nodes))

	for _, n := range nodes.Nodes {
		node := newNodeStats()
		node.ID = n.Desc.NodeID
		node.Ranges = uint(n.StoreStatuses[0].Desc.Capacity.RangeCount)

		metrics := n.StoreStatuses[0].Metrics
		node.Replicas = uint(metrics.Replicas)
		node.CapacityUsage = uint64(metrics.CapacityUsed)
		node.CapacityReserved = uint64(metrics.CapacityReserved)
		node.CapacityAvailable = uint64(metrics.CapacityAvailable)
		node.Capacity = uint64(metrics.Capacity)
		node.TotalSystemMemory, _ = strconv.ParseInt(n.TotalSystemMemory, 10, 64)
		node.QueriesPerSecond = n.StoreStatuses[0].Desc.Capacity.QueriesPerSecond
		node.Latency99 = uint64(n.Metrics.LivenessHeartbeatlatencyP99)
		node.ClockOffset = uint32(n.Metrics.ClockOffsetMeannanos)

		nodeStats = append(nodeStats, node)
	}

	return nodeStats
}

// getNodes returns a structure containing the response to a call to /_status/nodes
func (rc RoachConnector) getNodes() ([]*NodeStats, error) {
	resp, err := rc.get(rc.ip + ":" + rc.port + reqAllNodes)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Error("Closing response body failed. But who cares ?")
		}
	}()

	nodes := &AllNodes{}
	err = json.NewDecoder(resp.Body).Decode(nodes)
	if err != nil {
		log.Error("getNodes : Error in parsing json : ", err)
		return nil, err
	}
	if len(nodes.Nodes) <= 0 {
		return nil, errors.New("received positive response for 0 nodes")
	}
	log.Infof("Extracting info from %d nodes.", len(nodes.Nodes))
	return extractNodeInfo(nodes), nil
}

/*
func (rc RoachConnector) getRanges() (int, error) {
	resp, err := rc.get(rc.ip + ":" + rc.port + reqClusterRaft)
	if err != nil {
		if resp != nil {
			return 0, err
		}
		return 0, err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Error("Closing response body failed. But who cares ?")
		}
	}()

	raft := &ClusterRaft{}
	err = json.NewDecoder(resp.Body).Decode(raft)
	if err != nil {
		log.Error("getRanges : Error in parsing json : ", err)
	} else {
		//log.Infof("getNodes : Parsed json successfully : %#v", nodes)
	}
	//log.Infof("Length of nodes : %d", len(ranges))

	ranges := len(raft.Ranges)
	log.Infof("getRanges : counted %d ranges", ranges)

	return ranges, nil
}
*/

// getHealth sends a single request to the specified resource
func (rc RoachConnector) getHealth(fullAdd string) (status string, result []byte, err error) {
	resp, err := rc.get(fullAdd)
	if err != nil {
		if resp != nil {
			return resp.Status, nil, err
		}
		return "", nil, err
	}

	//log.Info("Response status: ", resp.Status)
	raw, _ := ioutil.ReadAll(resp.Body)

	//var h Health
	var test1 map[string]interface{}

	err = json.Unmarshal(raw, &test1)
	if err != nil {
		log.Error("Error in parsing json : ", err)
		return "0", nil, err
	}
	//log.Infof("Parsed json successfully : %#v", test1)

	//log.Info("Raw response ", raw)
	//log.Info("String response ", string(raw))
	//return resp.Status, string(raw), nil
	return resp.Status, raw, nil
}

/**
**
** Implement Connector on RoachConnector
**
 */

// Probe operates requests to the cluster
func (rc RoachConnector) Probe() (probe *dbmon.Probe, err error) {

	log.Info("Probing cluster ", rc.clusterID)

	// Get Number of nodes and their ID
	nodeStats, err := rc.getNodes()
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(nodeStats)
	if err != nil {
		return nil, err
	}

	probe = &dbmon.Probe{
		ClusterID:            rc.clusterID,
		Status:               "200",
		Timestamp:            time.Now().String(),
		Data:                 payload,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	log.Info("Probing finished, sending payload to collector.")
	return probe, nil
}

// Period returns the connector's period attribute
func (rc RoachConnector) Period() time.Duration {
	return rc.period
}
