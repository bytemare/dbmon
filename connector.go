package dbmon

import (
	"time"
)

// Connector interface
// Connector is a request abstraction layer to the dbmon collector.
// It will adapt generic queries to the specified database endpoint.
type Connector interface {
	// Probe asks the connector to operate a probe to the cluster
	Probe() (probe *Probe, err error)

	// Period returns a connector's period attribute
	Period() time.Duration
}
