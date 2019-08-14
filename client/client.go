// Client is a dbmon CLI client connecting to the server to pull data
// It then can display, store, or relay it to a monitoring tool.
package client

type client struct {
	endpoint string // Endpoint domain or IP to interrogate
}
