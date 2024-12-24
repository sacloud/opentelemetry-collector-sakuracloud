package sacloud

import (
	"context"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/packages-go/newsfeed"
)

type LoadBalancer struct {
	*iaas.LoadBalancer
	ZoneName string
}

type LoadBalancerClient interface {
	Find(ctx context.Context) ([]*LoadBalancer, error)
	Status(ctx context.Context, zone string, id types.ID) ([]*iaas.LoadBalancerStatus, error)
	MonitorNIC(ctx context.Context, zone string, id types.ID, end time.Time) (*iaas.MonitorInterfaceValue, error)
	MaintenanceInfo(infoURL string) (*newsfeed.FeedItem, error)
}

func getLoadBalancerClient(caller iaas.APICaller, zones []string) LoadBalancerClient {
	return &loadBalancerClient{
		client: iaas.NewLoadBalancerOp(caller),
		zones:  zones,
	}
}

type loadBalancerClient struct {
	client iaas.LoadBalancerAPI
	zones  []string
}

func (c *loadBalancerClient) find(ctx context.Context, zone string) ([]interface{}, error) {
	var results []interface{}
	res, err := c.client.Find(ctx, zone, &iaas.FindCondition{
		Count: 10000,
	})
	if err != nil {
		return results, err
	}
	for _, lb := range res.LoadBalancers {
		results = append(results, &LoadBalancer{
			LoadBalancer: lb,
			ZoneName:     zone,
		})
	}
	return results, err
}

func (c *loadBalancerClient) Find(ctx context.Context) ([]*LoadBalancer, error) {
	res, err := queryToZones(ctx, c.zones, c.find)
	if err != nil {
		return nil, err
	}
	var results []*LoadBalancer
	for _, s := range res {
		results = append(results, s.(*LoadBalancer))
	}
	return results, nil
}

func (c *loadBalancerClient) MonitorNIC(ctx context.Context, zone string, id types.ID, end time.Time) (*iaas.MonitorInterfaceValue, error) {
	mvs, err := c.client.MonitorInterface(ctx, zone, id, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorInterfaceValue(mvs.Values), nil
}

func (c *loadBalancerClient) Status(ctx context.Context, zone string, id types.ID) ([]*iaas.LoadBalancerStatus, error) {
	res, err := c.client.Status(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	return res.Status, nil
}

func (c *loadBalancerClient) MaintenanceInfo(infoURL string) (*newsfeed.FeedItem, error) {
	return newsfeed.GetByURL(infoURL)
}
