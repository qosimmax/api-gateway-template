// Package fakeapi contains an api api service.
package fakeapi

import (
	"api-gateway-template/app/pb/fakeapi"
	"api-gateway-template/config"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client holds the api service.
type Client struct {
	Endpoint string
	conn     *grpc.ClientConn
	service  fakeapi.FakeServiceClient
}

// Init initializes a new api service.
func (c *Client) Init(config *config.Config) error {
	// Set up a connection to the server.
	var err error
	c.conn, err = grpc.Dial(config.ExampleAPIEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		return fmt.Errorf("fake-api connection error: %w", err)
	}

	c.service = fakeapi.NewFakeServiceClient(c.conn)
	return nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
