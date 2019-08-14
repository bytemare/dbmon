// cockroachConnector implements the Connector interface
// It hols everything needed to connect to a CockroachDB cluster and retrieve data
package connectors

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	// Health enpoints
	roachUriHealth		=	"/health"
	//roachUriHealthReady		=	"/health?ready=1"
)


type RoachConnector struct {
	ip			string
	port		string
	period    	time.Duration // Period to repetitively interrogate the cluster
}

// New return
func NewConnector(ip string, port string) *RoachConnector {
	return &RoachConnector{
		ip:ip,
		port:port,
		period: 2 * time.Second,
	}
}

/**
**
** Implement Connector on RoachConnector
**
 */

// Request operates a request
func (rc RoachConnector) Request() (response string, err error) {

	log.Info("Sending request...")

	resp, err := http.Get(rc.ip + ":" + rc.port + roachUriHealth )
	if err != nil {
		log.Error(err)
		return "", err
	}
	defer resp.Body.Close()

	log.Info("Response status: ", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		response += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Error("Scanner err: ", err)
		return "", err
	}

	return response,nil
}

// Generic Information about the cluster
func (rc RoachConnector) GetIP() (ip string){
	return rc.ip
}

// Return the period for requests
func (rc RoachConnector) Period() (period time.Duration){
	return rc.period
}

// Returns a generic status about the cluster (OK, Not available, slow)
func (rc RoachConnector) Status() (status int, err error){
	return 200, nil
}