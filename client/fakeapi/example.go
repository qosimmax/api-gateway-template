package fakeapi

import (
	"api-gateway-template/app"
	"api-gateway-template/app/pb/fakeapi"
	"context"
)

func (c *Client) GetExampleData(ctx context.Context) (*app.Data, error) {
	_, err := c.service.GetFakeRequest(ctx, &fakeapi.FakeRequest{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
