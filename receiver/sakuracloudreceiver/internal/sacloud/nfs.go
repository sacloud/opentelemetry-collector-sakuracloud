package sacloud

import (
	"context"
	"time"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/helper/query"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/packages-go/newsfeed"
)

type NFS struct {
	*iaas.NFS
	Plan     *query.NFSPlanInfo
	PlanName string
	ZoneName string
}

type NFSClient interface {
	Find(ctx context.Context) ([]*NFS, error)
	MonitorFreeDiskSize(ctx context.Context, zone string, id types.ID, end time.Time) (*iaas.MonitorFreeDiskSizeValue, error)
	MonitorNIC(ctx context.Context, zone string, id types.ID, end time.Time) (*iaas.MonitorInterfaceValue, error)
	MaintenanceInfo(infoURL string) (*newsfeed.FeedItem, error)
}

func getNFSClient(caller iaas.APICaller, zones []string) NFSClient {
	return &nfsClient{
		noteOp: iaas.NewNoteOp(caller),
		nfsOp:  iaas.NewNFSOp(caller),
		zones:  zones,
	}
}

type nfsClient struct {
	noteOp iaas.NoteAPI
	nfsOp  iaas.NFSAPI
	zones  []string
}

func (c *nfsClient) find(ctx context.Context, zone string) ([]interface{}, error) {
	var results []interface{}
	res, err := c.nfsOp.Find(ctx, zone, &iaas.FindCondition{
		Count: 10000,
	})
	if err != nil {
		return results, err
	}
	for _, v := range res.NFS {
		planInfo, err := query.GetNFSPlanInfo(ctx, c.noteOp, v.PlanID)
		if err != nil {
			return nil, err
		}
		planName := ""
		switch planInfo.DiskPlanID {
		case types.NFSPlans.HDD:
			planName = "HDD"
		case types.NFSPlans.SSD:
			planName = "SSD"
		}
		results = append(results, &NFS{
			NFS:      v,
			PlanName: planName,
			Plan:     planInfo,
			ZoneName: zone,
		})
	}
	return results, err
}

func (c *nfsClient) Find(ctx context.Context) ([]*NFS, error) {
	res, err := queryToZones(ctx, c.zones, c.find)
	if err != nil {
		return nil, err
	}
	var results []*NFS
	for _, s := range res {
		results = append(results, s.(*NFS))
	}
	return results, nil
}

func (c *nfsClient) MonitorFreeDiskSize(ctx context.Context, zone string, id types.ID, end time.Time) (*iaas.MonitorFreeDiskSizeValue, error) {
	mvs, err := c.nfsOp.MonitorFreeDiskSize(ctx, zone, id, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorFreeDiskSizeValue(mvs.Values), nil
}

func (c *nfsClient) MonitorNIC(ctx context.Context, zone string, id types.ID, end time.Time) (*iaas.MonitorInterfaceValue, error) {
	mvs, err := c.nfsOp.MonitorInterface(ctx, zone, id, monitorCondition(end))
	if err != nil {
		return nil, err
	}
	return monitorInterfaceValue(mvs.Values), nil
}

func (c *nfsClient) MaintenanceInfo(infoURL string) (*newsfeed.FeedItem, error) {
	return newsfeed.GetByURL(infoURL)
}
