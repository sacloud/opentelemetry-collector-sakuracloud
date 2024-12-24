package sacloud

import (
	"context"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

type ESMEClient interface {
	Find(ctx context.Context) ([]*iaas.ESME, error)
	Logs(ctx context.Context, esmeID types.ID) ([]*iaas.ESMELogs, error)
}

func getESMEClient(caller iaas.APICaller) ESMEClient {
	return &esmeClient{
		caller: caller,
	}
}

type esmeClient struct {
	caller iaas.APICaller
}

func (c *esmeClient) Find(ctx context.Context) ([]*iaas.ESME, error) {
	client := iaas.NewESMEOp(c.caller)
	searched, err := client.Find(ctx, &iaas.FindCondition{})
	if err != nil {
		return nil, err
	}
	return searched.ESME, nil
}

func (c *esmeClient) Logs(ctx context.Context, esmeID types.ID) ([]*iaas.ESMELogs, error) {
	client := iaas.NewESMEOp(c.caller)
	return client.Logs(ctx, esmeID)
}
