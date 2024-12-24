package sacloud

import (
	"context"

	"github.com/sacloud/webaccel-api-go"
)

// WebAccelClient calls SakuraCloud webAccel API
type WebAccelClient interface {
	Find(ctx context.Context) ([]*webaccel.Site, error)
	Usage(ctx context.Context) (*webaccel.MonthlyUsageResults, error)
}

func getWebAccelClient(caller webaccel.APICaller) WebAccelClient {
	return &webAccelClient{
		client: webaccel.NewOp(caller),
	}
}

type webAccelClient struct {
	client webaccel.API
}

func (c *webAccelClient) Find(ctx context.Context) ([]*webaccel.Site, error) {
	res, err := c.client.List(ctx)
	if err != nil {
		return nil, err
	}
	return res.Sites, nil
}

func (c *webAccelClient) Usage(ctx context.Context) (*webaccel.MonthlyUsageResults, error) {
	return c.client.MonthlyUsage(ctx, "")
}
