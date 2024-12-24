package sacloud

import (
	"context"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

type LocalRouterClient interface {
	Find(ctx context.Context) ([]*iaas.LocalRouter, error)
	Health(ctx context.Context, id types.ID) (*iaas.LocalRouterHealth, error)
	Monitor(ctx context.Context, id types.ID, end time.Time) (*iaas.MonitorLocalRouterValue, error)
}

func getLocalRouterClient(caller iaas.APICaller) LocalRouterClient {
	return &localRouterClient{
		client: iaas.NewLocalRouterOp(caller),
	}
}

type localRouterClient struct {
	client iaas.LocalRouterAPI
}

func (c *localRouterClient) Find(ctx context.Context) ([]*iaas.LocalRouter, error) {
	var results []*iaas.LocalRouter
	res, err := c.client.Find(ctx, &iaas.FindCondition{
		Count: 10000,
	})
	if err != nil {
		return results, err
	}
	return res.LocalRouters, nil
}

func (c *localRouterClient) Health(ctx context.Context, id types.ID) (*iaas.LocalRouterHealth, error) {
	return c.client.HealthStatus(ctx, id)
}

func (c *localRouterClient) Monitor(ctx context.Context, id types.ID, end time.Time) (*iaas.MonitorLocalRouterValue, error) {
	mvs, err := c.client.MonitorLocalRouter(ctx, id, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorLocalRouterValue(mvs.Values), nil
}
