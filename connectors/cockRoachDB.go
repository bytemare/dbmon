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

	// TODO : do not use this in production, DEMO ONLY
	resp, err := http.Get(rc.ip + ":" + rc.port + roachUriHealth)
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
