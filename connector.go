package dbmon

import (
	"time"
)

// Connector interface
// Connector is a request abstraction layer to the dbmon collector.
// It will adapt generic queries to the specified database endpoint.
type Connector interface {
/*
	// Metadata about a connector for the collector
	Id() (id string)
	Ip() (ip string)
	Port() (port string)
	Status() (status string) // Responsive/Online, Pending, Unauthorised, Unreachable, Stopped/Offline
	Stop() // Tells the connector to stop and free resources
	// TODO : maybe use closures to repetitively work on same object ?


*/
	// Probe asks the connector to operate a probe to the cluster
	Probe() (probe *Probe, err error)

	// Period returns a connector's period attribute
	Period() time.Duration
}
