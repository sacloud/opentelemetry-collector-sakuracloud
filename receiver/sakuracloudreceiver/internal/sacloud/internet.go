package sacloud

import (
	"context"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

type Internet struct {
	*iaas.Internet
	ZoneName string
}

type InternetClient interface {
	Find(ctx context.Context) ([]*Internet, error)
	MonitorTraffic(ctx context.Context, zone string, internetID types.ID, end time.Time) (*iaas.MonitorRouterValue, error)
}

func getInternetClient(caller iaas.APICaller, zones []string) InternetClient {
	return &internetClient{
		client: iaas.NewInternetOp(caller),
		zones:  zones,
	}
}

type internetClient struct {
	client iaas.InternetAPI
	zones  []string
}

func (c *internetClient) find(ctx context.Context, zone string) ([]interface{}, error) {
	var results []interface{}
	res, err := c.client.Find(ctx, zone, &iaas.FindCondition{
		Count: 10000,
	})
	if err != nil {
		return results, err
	}
	for _, router := range res.Internet {
		results = append(results, &Internet{
			Internet: router,
			ZoneName: zone,
		})
	}
	return results, err
}

func (c *internetClient) Find(ctx context.Context) ([]*Internet, error) {
	res, err := queryToZones(ctx, c.zones, c.find)
	if err != nil {
		return nil, err
	}
	var results []*Internet
	for _, s := range res {
		results = append(results, s.(*Internet))
	}
	return results, nil
}

func (c *internetClient) MonitorTraffic(ctx context.Context, zone string, internetID types.ID, end time.Time) (*iaas.MonitorRouterValue, error) {
	mvs, err := c.client.Monitor(ctx, zone, internetID, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorRouterValue(mvs.Values), nil
}
