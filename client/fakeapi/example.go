package fakeapi

import (
	"api-gateway-template/app"
	"api-gateway-template/app/pb/fakeapi"
	"context"
	"time"
)

func (c *Client) GetExampleData(ctx context.Context) (*app.Data, error) {
	resp, err := c.service.GetFakeRequest(ctx, &fakeapi.FakeRequest{})
	if err != nil {
		return nil, err
	}

	return &app.Data{
		Message: resp.GetMessage(),
		Date:    time.Now(),
	}, nil
}
