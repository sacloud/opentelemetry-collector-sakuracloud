package sacloud

import (
	"context"

	"github.com/sacloud/iaas-api-go"
)

// ZoneClient calls SakuraCloud zone API
type ZoneClient interface {
	Find(ctx context.Context) ([]*iaas.Zone, error)
}

func getZoneClient(caller iaas.APICaller) ZoneClient {
	return &zoneClient{
		client: iaas.NewZoneOp(caller),
	}
}

type zoneClient struct {
	client iaas.ZoneAPI
}

func (c *zoneClient) Find(ctx context.Context) ([]*iaas.Zone, error) {
	res, err := c.client.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	return res.Zones, nil
}
