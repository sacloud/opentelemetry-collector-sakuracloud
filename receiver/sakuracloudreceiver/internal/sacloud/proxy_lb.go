package sacloud

import (
	"context"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

type ProxyLBClient interface {
	Find(ctx context.Context) ([]*iaas.ProxyLB, error)
	GetCertificate(ctx context.Context, id types.ID) (*iaas.ProxyLBCertificates, error)
	Monitor(ctx context.Context, id types.ID, end time.Time) (*iaas.MonitorConnectionValue, error)
}

func getProxyLBClient(caller iaas.APICaller) ProxyLBClient {
	return &proxyLBClient{
		client: iaas.NewProxyLBOp(caller),
	}
}

type proxyLBClient struct {
	client iaas.ProxyLBAPI
}

func (c *proxyLBClient) Find(ctx context.Context) ([]*iaas.ProxyLB, error) {
	var results []*iaas.ProxyLB
	res, err := c.client.Find(ctx, &iaas.FindCondition{
		Count: 10000,
	})
	if err != nil {
		return results, err
	}
	return res.ProxyLBs, nil
}

func (c *proxyLBClient) GetCertificate(ctx context.Context, id types.ID) (*iaas.ProxyLBCertificates, error) {
	return c.client.GetCertificates(ctx, id)
}

func (c *proxyLBClient) Monitor(ctx context.Context, id types.ID, end time.Time) (*iaas.MonitorConnectionValue, error) {
	mvs, err := c.client.MonitorConnection(ctx, id, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorConnectionValue(mvs.Values), nil
}
