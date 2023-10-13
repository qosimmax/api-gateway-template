// Package api contains an api api client.
package api

import "api-gateway-template/config"

// Client holds the api client.
type Client struct {
	Endpoint string
}

// Init initializes a new api client.
func (c *Client) Init(config *config.Config) error {
	return nil
}
