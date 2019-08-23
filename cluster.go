package dbmon

// Cluster // Cluster represents a registered Database Cluster to a dbmon instance.
// It holds identity information about a Cluster and a corresponding Connector.
type Cluster struct {
	id        string    // Identification/Name for the target cluster
	connector Connector // Connector enabling communicating with the cluster
}

// NewCluster returns a new initialised Cluster struct
func NewCluster(id string, c Connector) *Cluster {
	return &Cluster{
		id:        id,
		connector: c,
	}
}
