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
	sink       chan *Probe   // Channel agents send their data to
	serverChan chan<- *Probe // Channel to send data to
	newCluster chan Cluster        // Channel to receive new requests for cluster registration

	// Synchronisation for closing
	sync chan struct{} // Channel used to instruct Collector to stop
	wait sync.WaitGroup
}

// NewCollector initialises and returns a new Collector struct
func NewCollector(serverChan chan<- *Probe) *Collector {
	return &Collector{
		clusters:   make(map[string]*Cluster),
		agents:     make(map[*Cluster]*agent),
		sink:       make(chan *Probe),
		serverChan: serverChan,
		newCluster: make(chan Cluster, 1000),
		sync:       make(chan struct{}),
		wait:       sync.WaitGroup{},
	}
}

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

	agent := newAgent(cluster, c.sink, wg)

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

// RegisterNewCluster registers a new cluster to the collector, spawning a new agent for it
func (c *Collector) RegisterNewCluster(cluster *Cluster) {
	// todo : if collector is not running, add it directly. use a running flag
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

		// Todo : use a message type method,
		//  by using a single channel to control collector : add cluster, remove cluster, etc.

		case <-c.sync:
			log.Info("Collector received the stop signal.")
			c.stopAll()
			close(c.sync)
			close(c.sink)
			break collector

		case probe := <-c.sink:
			log.Info("Collector got new probe !")
			// Send to server
			// TODO : if changing architecture, use here a connection to the server
			c.serverChan <- probe
			log.Info("Collector : probe sent to server.")

		case cluster := <-c.newCluster:
			if err := c.newAgent(&cluster, &c.wait); err != nil {
				log.WithFields(log.Fields{
					"cluster":   cluster.id,
					"timestamp": time.Now().Format(timeLayout),
					"raw data":  cluster,
				}).Error(err)
			}
		}
	}
}

// Stop stops the Collector
func (c *Collector) Stop() {
	c.sync <- struct{}{}
	c.wait.Wait()
}
