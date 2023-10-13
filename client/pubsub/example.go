package pubsub

import (
	"api-gateway-template/app"
	"context"
	"encoding/json"
	"fmt"
)

// NotifyExampleData is used to publish example data to pubsub.
func (c *Client) NotifyExampleData(ctx context.Context, exampleData app.Data) error {
	data, err := json.Marshal(exampleData)
	if err != nil {
		return fmt.Errorf("error marshalling example data to send to pubsub: %w", err)
	}
	err = c.send(ctx, "whatever-topic-name", data)
	if err != nil {
		return fmt.Errorf("error sending example data message to pubsub: %w", err)
	}
	return nil
}
