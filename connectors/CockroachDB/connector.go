// cockroachConnector implements the Connector interface
// It hols everything needed to connect to a CockroachDB cluster and retrieve data
package cockroachdb

/**
TODO : The json API is subject to change, so we're sending the raw json response, for now.
TODO : Maybe using the cockroachdb command could help : but how would we do this with a remote cluster ?
*/

import (
	"github.com/bytemare/dbmon/dbmon"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
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
		period:    2 * time.Second,
		timeout:   roachTimeout,
		netClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

/**
**
** Implement Connector on RoachConnector
**
 */

// sendRequest sends a single request to the specified ressource
func (rc RoachConnector) sendRequest(fullAdd string) (status string, result []byte, err error) {
	resp, err := rc.netClient.Get(fullAdd)
	if err != nil {
		log.Error(err)
		return resp.Status, nil, err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Error("Closing response body failed. But who cares ?")
		}
	}()

	log.Info("Response status: ", resp.Status)
	raw, _ := ioutil.ReadAll(resp.Body)
	//log.Info("Raw response ", raw)
	//log.Info("String response ", string(raw))
	//return resp.Status, string(raw), nil
	return resp.Status, raw, nil
}

// Probe operates requests to the cluster
func (rc RoachConnector) Probe() (probe *dbmon.Probe, err error) {

	log.Info("Sending request to Cluster ", rc.clusterId)
	status, res, err := rc.sendRequest(rc.ip + ":" + rc.port + reqHealth)
	if err != nil {
		return nil, err
	}

	probe = &dbmon.Probe{
		ClusterID:            rc.clusterId,
		Status:               status,
		Timestamp:            time.Now().String(),
		Len:                  []int32{int32(len(res))},
		Data:                 [][]byte{res},
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
