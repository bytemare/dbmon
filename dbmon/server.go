// dbmon is a server daemon with a gRPC API endpoint that serves information about a database cluster
package dbmon

// DBMon
type DBMon struct {
	data   map[*Cluster]report // Cache to hold data to serve
	source <-chan [2]string    // Channel through which data is arriving
}

// report holds a cluster's accumulated information
type report struct {
	clusterID *Cluster
	data      []string
}

// NewDBMon initialises and returns a new DBMon struct
func NewDBMon() (*DBMon, error) {
	return nil, nil
}

// Start starts the DBMon server
func (mon *DBMon) Start() error {
	return nil
}

// Stop stops the DBMon server
func (mon *DBMon) Stop() error {
	return nil
}
