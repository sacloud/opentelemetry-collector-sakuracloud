package sacloud

import (
	"context"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

type SIMClient interface {
	Find(ctx context.Context) ([]*iaas.SIM, error)
	GetNetworkOperatorConfig(ctx context.Context, id types.ID) ([]*iaas.SIMNetworkOperatorConfig, error)
	MonitorTraffic(ctx context.Context, id types.ID, end time.Time) (*iaas.MonitorLinkValue, error)
}

func getSIMClient(caller iaas.APICaller) SIMClient {
	return &simClient{
		client: iaas.NewSIMOp(caller),
	}
}

type simClient struct {
	client iaas.SIMAPI
}

func (c *simClient) Find(ctx context.Context) ([]*iaas.SIM, error) {
	var results []*iaas.SIM
	res, err := c.client.Find(ctx, &iaas.FindCondition{
		Include: []string{"*", "Status.sim"},
		Count:   10000,
	})
	if err != nil {
		return results, err
	}
	return res.SIMs, nil
}

func (c *simClient) GetNetworkOperatorConfig(ctx context.Context, id types.ID) ([]*iaas.SIMNetworkOperatorConfig, error) {
	return c.client.GetNetworkOperator(ctx, id)
}

func (c *simClient) MonitorTraffic(ctx context.Context, id types.ID, end time.Time) (*iaas.MonitorLinkValue, error) {
	mvs, err := c.client.MonitorSIM(ctx, id, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorLinkValue(mvs.Values), nil
}
