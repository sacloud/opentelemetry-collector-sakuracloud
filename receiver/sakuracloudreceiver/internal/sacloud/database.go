package sacloud

import (
	"context"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/packages-go/newsfeed"
)

type Database struct {
	*iaas.Database
	ZoneName string
}

type DatabaseClient interface {
	Find(ctx context.Context) ([]*Database, error)
	MonitorDatabase(ctx context.Context, zone string, diskID types.ID, end time.Time) (*iaas.MonitorDatabaseValue, error)
	MonitorCPU(ctx context.Context, zone string, databaseID types.ID, end time.Time) (*iaas.MonitorCPUTimeValue, error)
	MonitorNIC(ctx context.Context, zone string, databaseID types.ID, end time.Time) (*iaas.MonitorInterfaceValue, error)
	MonitorDisk(ctx context.Context, zone string, databaseID types.ID, end time.Time) (*iaas.MonitorDiskValue, error)
	MaintenanceInfo(infoURL string) (*newsfeed.FeedItem, error)
}

func getDatabaseClient(caller iaas.APICaller, zones []string) DatabaseClient {
	return &databaseClient{
		client: iaas.NewDatabaseOp(caller),
		zones:  zones,
	}
}

type databaseClient struct {
	client iaas.DatabaseAPI
	zones  []string
}

func (c *databaseClient) find(ctx context.Context, zone string) ([]interface{}, error) {
	var results []interface{}
	res, err := c.client.Find(ctx, zone, &iaas.FindCondition{
		Count: 10000,
	})
	if err != nil {
		return results, err
	}
	for _, db := range res.Databases {
		results = append(results, &Database{
			Database: db,
			ZoneName: zone,
		})
	}
	return results, err
}

func (c *databaseClient) Find(ctx context.Context) ([]*Database, error) {
	res, err := queryToZones(ctx, c.zones, c.find)
	if err != nil {
		return nil, err
	}
	var results []*Database
	for _, s := range res {
		results = append(results, s.(*Database))
	}
	return results, nil
}

func (c *databaseClient) MonitorDatabase(ctx context.Context, zone string, databaseID types.ID, end time.Time) (*iaas.MonitorDatabaseValue, error) {
	mvs, err := c.client.MonitorDatabase(ctx, zone, databaseID, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorDatabaseValue(mvs.Values), nil
}

func (c *databaseClient) MonitorCPU(ctx context.Context, zone string, databaseID types.ID, end time.Time) (*iaas.MonitorCPUTimeValue, error) {
	mvs, err := c.client.MonitorCPU(ctx, zone, databaseID, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorCPUTimeValue(mvs.Values), nil
}

func (c *databaseClient) MonitorDisk(ctx context.Context, zone string, databaseID types.ID, end time.Time) (*iaas.MonitorDiskValue, error) {
	mvs, err := c.client.MonitorDisk(ctx, zone, databaseID, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorDiskValue(mvs.Values), nil
}

func (c *databaseClient) MonitorNIC(ctx context.Context, zone string, databaseID types.ID, end time.Time) (*iaas.MonitorInterfaceValue, error) {
	mvs, err := c.client.MonitorInterface(ctx, zone, databaseID, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorInterfaceValue(mvs.Values), nil
}

func (c *databaseClient) MaintenanceInfo(infoURL string) (*newsfeed.FeedItem, error) {
	return newsfeed.GetByURL(infoURL)
}
