// cockroachConnector implements the Connector interface
// It hols everything needed to connect to a CockroachDB cluster and retrieve data
package connectors

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	// Health enpoints
	roachUriHealth = "/health"
	//roachUriHealthReady		=	"/health?ready=1"
	roachTimeout = 10 * time.Second
)

type RoachConnector struct {
	ip     string
	port   string
	period time.Duration // Period to repetitively interrogate the cluster
}

// New returns a new initialised connector for a CockRoachDB endpoint
func NewConnector(ip string, port string) *RoachConnector {
	log.Warn("Warning : DEMO ! CockroachDB should not use vanilla http.Client for production use.")
	return &RoachConnector{
		ip:     ip,
		port:   port,
		period: 2 * time.Second,
	}
}

/**
**
** Implement Connector on RoachConnector
**
 */

type response struct {
	body string
}

func decode(resp *http.Response, result interface{}) error {
	return json.NewDecoder(resp.Body).Decode(result)
}

// Request operates a request
func (rc RoachConnector) Request() (result string, err error) {

	log.Info("Sending request...")

	// Set the timeout on client
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := netClient.Get(rc.ip + ":" + rc.port)
	if err != nil {
		log.Error(err)
		return "", err
	}
	defer resp.Body.Close()

	log.Info("Response status: ", resp.Status)

	decoded := new(response)
	err = decode(resp, decoded)
	if err != nil {
		log.Error("Error in decoding json response : ", err)
		return "", err
	}

	log.Info("Decoded result : ", decoded.body)

	return decoded.body, nil
}

// Generic Information about the cluster
func (rc RoachConnector) GetIP() (ip string) {
	return rc.ip
}

// Return the period for requests
func (rc RoachConnector) Period() (period time.Duration) {
	return rc.period
}

// Returns a generic status about the cluster (OK, Not available, slow)
func (rc RoachConnector) Status() (status int, err error) {
	return 200, nil
}

/**
**
** Json objects : TODO : Is there a better way of doing this ?
**
 */


/*
Top structures specific to a request
 */

// Health is a JSON object returned from a /health request
type Health struct {
	NodeID     int        `json:"nodeId"`
	Address    Address    `json:"address"`
	BuildInfo  BuildInfo  `json:"buildInfo"`
	SystemInfo SystemInfo `json:"systemInfo"`
}

//
type AutoGenerated struct {
	Events []Events `json:"events"`
}






/*

 */

type Address struct {
	NetworkField string `json:"networkField"`
	AddressField string `json:"addressField"`
}
type BuildInfo struct {
	GoVersion       string      `json:"goVersion"`
	Tag             string      `json:"tag"`
	Time            string      `json:"time"`
	Revision        string      `json:"revision"`
	CgoCompiler     string      `json:"cgoCompiler"`
	CgoTargetTriple string      `json:"cgoTargetTriple"`
	Platform        string      `json:"platform"`
	Distribution    string      `json:"distribution"`
	Type            string      `json:"type"`
	Channel         string      `json:"channel"`
	EnvChannel      string      `json:"envChannel"`
	Dependencies    interface{} `json:"dependencies"`
}
type SystemInfo struct {
	SystemInfo string `json:"systemInfo"`
	KernelInfo string `json:"kernelInfo"`
}


type Replicas struct {
	NodeID    int `json:"nodeId"`
	StoreID   int `json:"storeId"`
	ReplicaID int `json:"replicaId"`
}
type UpdatedDesc struct {
	RangeID       string      `json:"rangeId"`
	StartKey      interface{} `json:"startKey"`
	EndKey        interface{} `json:"endKey"`
	Replicas      []Replicas  `json:"replicas"`
	NextReplicaID int         `json:"nextReplicaId"`
	Generation    string      `json:"generation"`
}
type NewDesc struct {
	RangeID       string      `json:"rangeId"`
	StartKey      interface{} `json:"startKey"`
	EndKey        interface{} `json:"endKey"`
	Replicas      []Replicas  `json:"replicas"`
	NextReplicaID int         `json:"nextReplicaId"`
	Generation    interface{} `json:"generation"`
}
type Info struct {
	UpdatedDesc    UpdatedDesc `json:"updatedDesc"`
	NewDesc        NewDesc     `json:"newDesc"`
	RemovedDesc    interface{} `json:"removedDesc"`
	AddedReplica   interface{} `json:"addedReplica"`
	RemovedReplica interface{} `json:"removedReplica"`
	Reason         string      `json:"reason"`
	Details        string      `json:"details"`
}
type Event struct {
	Timestamp    time.Time `json:"timestamp"`
	RangeID      string    `json:"rangeId"`
	StoreID      int       `json:"storeId"`
	EventType    int       `json:"eventType"`
	OtherRangeID string    `json:"otherRangeId"`
	Info         Info      `json:"info"`
}
type PrettyInfo struct {
	UpdatedDesc    string `json:"updatedDesc"`
	NewDesc        string `json:"newDesc"`
	AddedReplica   string `json:"addedReplica"`
	RemovedReplica string `json:"removedReplica"`
	Reason         string `json:"reason"`
	Details        string `json:"details"`
}
type Events struct {
	Event      Event      `json:"event"`
	PrettyInfo PrettyInfo `json:"prettyInfo"`
}