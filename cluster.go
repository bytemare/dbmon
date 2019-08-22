package dbmon

// Cluster // Cluster represents a registered Database Cluster to a dbmon instance.
// It holds identity information about a Cluster and a corresponding Connector.
type Cluster struct {
	id        string    // Identification/Name for the target cluster
	connector Connector // Connector enabling communicating with the cluster
}

// request defines a possible request for a cluster's API
type request struct {
	id      string   // Request identification
	target  *Cluster // Target cluster of the request
	request string   // Request to send to the cluster's API
}

// NewCluster returns a new initialised Cluster struct
func NewCluster(id string, c Connector) *Cluster {
	return &Cluster{
		id:        id,
		connector: c,
	}
}
