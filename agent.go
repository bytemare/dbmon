package dbmon

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

// agent is dedicated routine to handle connections to a cluster
type agent struct {
	cluster  *Cluster        // Target cluster
	requests []request       // Map to a cluster a slice of possible requests
	sink     chan<- *Probe   // Channel to send collected data to
	sync     chan struct{}   // Channel to be told to stop
	wg       *sync.WaitGroup // Synchronisation
}

func newAgent(cluster *Cluster, sink chan<- *Probe, wg *sync.WaitGroup) *agent {
	return &agent{
		cluster:  cluster,
		requests: []request{},
		sink:     sink,
		sync:     make(chan struct{}),
		wg:       wg,
	}
}

// run puts an agent in working mode, requesting indefinitely the associated cluster until being told to stop
func (a *agent) run() {
	a.wg.Add(1)
	defer a.wg.Done()

	ticker := time.NewTicker(a.cluster.connector.Period())

	log.Info("Agent starting for cluster ", a.cluster.id)

agent:
	for {
		select {

		case <-a.sync:
			log.WithFields(log.Fields{
				"cluster":   a.cluster.id,
				"timestamp": time.Now().Format(timeLayout),
			}).Info("Agent for cluster is now stopping.")
			// todo : verify if there aren't any pending connections or things to free
			close(a.sync)
			break agent

		case t := <-ticker.C:
			log.WithFields(log.Fields{
				"cluster":   a.cluster.id,
				"timestamp": t.Format(timeLayout),
			}).Info("Agent operates a probe.")

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

	probe, err := a.cluster.connector.Probe()
	if err != nil {
		return err
	}

	// Send data
	a.sink <- probe

	log.Info("Agent sent probe to collector.")

	return nil
}

// stop tells an agent to stop
func (a *agent) stop() {
	a.sync <- struct{}{}
}
