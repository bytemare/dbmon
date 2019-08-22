// cockroachConnector implements the Connector interface
// It hols everything needed to connect to a CockroachDB cluster and retrieve data
package cockroachdb

/**
TODO : The json API is subject to change, so we're sending the raw json response, for now.
TODO : Maybe using the cockroachdb command could help : but how would we do this with a remote cluster ?
*/

import (
	"encoding/json"
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
	reqHealth         = "/health"
	reqHealthReady    = "/health?ready=1"
	reqSingleNode     = "/_status/nodes/[node_id]"
	reqAllNodes       = "/_status/nodes"
	reqSingleHotrange = "_status/hotranges?node_id=[node_id]"
	reqAllHotRanges   = "/_status/hotranges"
	reqClusterRaft    = "/_status/raft"
	reqClusterRange   = "/_status/range/1"
)

type RoachConnector struct {
	clusterId string
	ip        string
	port      string
	period    time.Duration // Period to repetitively interrogate the cluster
	timeout   time.Duration // Timeout per request
	netClient *http.Client
}

// New returns a new initialised connector for a CockRoachDB endpoint
func NewConnector(id string, ip string, port string) *RoachConnector {
	log.Warn("Warning : DEMO ! CockroachDB should not use vanilla http.Client for production use.")
	return &RoachConnector{
		clusterId: id,
		ip:        ip,
		port:      port,
		period:    roachPeriod,
		timeout:   roachTimeout,
		netClient: &http.Client{
			Timeout: roachTimeout,
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

/**
**
** Implement Connector on RoachConnector
**
 */

// getNodes returns a structure containing the response to a call to /_status/nodes
func (rc RoachConnector) getNodes() ([]*NodeStats, error) {
	resp, err := rc.get(rc.ip + ":" + rc.port + reqAllNodes)
	if err != nil {
		if resp != nil {
			return nil, err
		}
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Error("Closing response body failed. But who cares ?")
		}
	}()

	nodes := &AllNodes{}
	/*raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error in reading json : ", err)
	} else {
		//log.Infof("Read json successfully : %#v", raw)
	}

	*/
	//err = json.Unmarshal(raw, &nodes)
	err = json.NewDecoder(resp.Body).Decode(nodes)
	if err != nil {
		log.Error("getNodes : Error in parsing json : ", err)
		return nil, err
	}

	nodeStats := make([]*NodeStats, 0, len(nodes.Nodes))

	for _, n := range nodes.Nodes {
		node := &NodeStats{
			ID:                0,
			Up:                false,
			Init:              false,
			Out:               false,
			Dead:              false,
			Ranges:            0,
			CapacityUsage:     0,
			CapacityReserved:  0,
			CapacityAvailable: 0,
			Capacity:          0,
			TotalSystemMemory: 0,
			MemoryUsage:       0,
			MemoryReserved:    0,
			MemoryAvailable:   0,
			QueriesPerSecond:  0,
			Latency99:         0,
			ClockOffset:       0,
		}
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
		//log.Infof("Node Info : %#v", node)
	}

	return nodeStats, nil
}

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
		//log.Error("Error in parsing json : ", err)
	}
	//log.Infof("Parsed json successfully : %#v", test1)

	//test2 := dyno.ConvertMapI2MapS(test1)
	//log.Infof("Remapped json : %#v", test2)

	//log.Info("Raw response ", raw)
	//log.Info("String response ", string(raw))
	//return resp.Status, string(raw), nil
	return resp.Status, raw, nil
}

// Probe operates requests to the cluster
func (rc RoachConnector) Probe() (probe *dbmon.Probe, err error) {

	log.Info("Sending request to Cluster ", rc.clusterId)
	/*status, res, err := rc.getHealth(rc.ip + ":" + rc.port + reqHealth)
	if err != nil {
		return nil, err
	}

	*/

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
		ClusterID:            rc.clusterId,
		Status:               "200",
		Timestamp:            time.Now().String(),
		Data:                 payload,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	return probe, nil
}

// Period returns the connector's period attribute
func (rc RoachConnector) Period() time.Duration {
	return rc.period
}
