// Collector is a service that interrogates database clusters to gather information
package dbmon

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const timeLayout = "2006-01-02 15:04:05:1234"

// Collector struct holds a list of agents to registered clusters
type Collector struct {
	clusters   map[string]*Cluster // List of currently registered clusters
	agents     map[*Cluster]*agent // List of running agents
	sink       chan [2]string      // Channel agents send their data to
	serverChan chan<- [2]string    // Channel to send data to
	newCluster chan Cluster      // Channel to receive new requests for cluster registration

	// Synchronisation for closing
	sync		chan struct{}     // Channel used to instruct Collector to stop
	wait		sync.WaitGroup
}

// agent is dedicated routine to handle connections to a cluster
type agent struct {
	cluster  *Cluster         // Target cluster
	requests []request        // Map to a cluster a slice of possible requests
	sink     chan<- [2]string // Channel to send collected data to
	sync     chan struct{}    // Channel to be told to stop
	wg       *sync.WaitGroup  // Synchronisation
}

// run puts an agent in working mode, requesting indefinitely the associated cluster until being told to stop
func (a *agent) run() {
	a.wg.Add(1)
	defer a.wg.Done()

	ticker := time.NewTicker(a.cluster.connector.Period())

agent:
	for {
		select {

		case <-a.sync:
			log.WithFields(log.Fields{
				"cluster":   a.cluster.id,
				"timestamp": time.Now().Format(timeLayout),
			}).Info("Agent for cluster is now stopping.")
			break agent

		case t := <-ticker.C:
			log.WithFields(log.Fields{
				"cluster":   a.cluster.id,
				"timestamp": t.Format(timeLayout),
			}).Info("Sending request to cluster.")

			if err := a.collect(); err != nil {
				log.WithFields(log.Fields{
					"cluster":   a.cluster.id,
					"timestamp": time.Now().Format(timeLayout),
					"request":   a.requests,
					"error":     err,
				}).Error("Error in collecting data from cluster.")
			}
		}
	}
	ticker.Stop()
}

// collect
func (a *agent) collect() error {

	data, err := a.cluster.connector.Request()
	if err != nil {
		log.Errorf("Agent encountered an error : %s", err)
	}

	// Send data
	a.sink <- [2]string{data, time.Now().Format(timeLayout)}

	return nil
}

// stop tells an agent to stop
func (a *agent) stop() {
	a.sync <- struct{}{}
}

/*
**
**
**	Collector
**
**
 */

// newAgent initialises and runs a new agent
func (c *Collector) newAgent(cluster *Cluster, wg *sync.WaitGroup) error {

	// Check whether the cluster is not already registered
	if _, inc := c.clusters[cluster.id]; inc == true {
		return errors.New("could not create new agent. An agent for this cluster already exists")
	}
	// The following tests should never succeed, but still
	if _, ina := c.agents[cluster]; ina == true {
		return errors.New("THIS SHOULD NOT HAPPEN. could not create new agent. An agent for this cluster already exists")
	}

	agent := &agent{
		cluster:  cluster,
		requests: []request{},
		sink:     c.sink,
		sync:     make(chan struct{}),
		wg:       wg,
	}

	c.clusters[cluster.id] = cluster
	c.agents[cluster] = agent

	go agent.run()

	return nil
}

// stopAgent tells the agent in charge for the cluster to stop
func (c *Collector) stopAgent(clusterID string) {
	cluster := c.clusters[clusterID]
	c.agents[cluster].stop()
	delete(c.agents, cluster)
	delete(c.clusters, clusterID)
}

// stopAll tells all agents to stop
func (c *Collector) stopAll() {
	for id := range c.clusters {
		c.stopAgent(id)
	}
}

// NewCollector initialises and returns a new Collector struct
func NewCollector(serverChan chan<- [2]string) *Collector {
	return &Collector{
		clusters:   make(map[string]*Cluster),
		agents:     make(map[*Cluster]*agent),
		sink:       make(chan [2]string),
		serverChan: serverChan,
		newCluster: make(chan Cluster),
		sync:       make(chan struct{}),
		wait:       sync.WaitGroup{},
	}
}

// RegisterNewCluster registers a new cluster to the collector, spawning a new agent for it
func (c *Collector) RegisterNewCluster(cluster *Cluster) {
		c.newCluster <- *cluster
}

// UnregisterCluster stops the corresponding agent to the cluster and deletes its entry in the register
func (c *Collector) UnregisterCluster(cluster *Cluster) {
	//todo
	log.Panic("OMG: UnregisterCluster IS NOT IMPLEMENTED !!!")
}

// Start starts the Collector
func (c *Collector) Start() {
collector:
	for {
		select {

		case <-c.sync:
			log.Info("Collector received the stop signal.")
			c.stopAll()
			break collector

		case data := <-c.sink:
			// Send to server
			c.serverChan <- data

		case cluster := <-c.newCluster:
			if err := c.newAgent(&cluster, &c.wait); err != nil {
				log.WithFields(log.Fields{
					"cluster":   cluster.id,
					"timestamp": time.Now().Format(timeLayout),
					"raw data":  cluster,
				}).Error(err)

		// Todo : use a message type method, by using a single channel to control collector : add cluster, remove cluster, etc.
			}
		}
	}
}

// Stop stops the Collector
func (c *Collector) Stop() {
	c.sync <- struct{}{}
	c.wait.Wait()
}
