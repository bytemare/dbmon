// Connector is a request abstraction layer to the dbmon collector.
// It will adapt generic queries to the specified database endpoint.
package dbmon

import "time"

// Connector interface
type Connector interface {
	Request() (response string, err error)
	//SetRequest(req string) (err error)

	// Generic Information about the cluster
	GetIP() (ip string)
	// Return the period for requests
	Period() (period time.Duration)
	// Returns a generic status about the cluster (OK, Not available, slow)
	Status() (status int, err error)
}
