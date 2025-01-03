package sacloud

import (
	"context"
	"testing"

	"github.com/sacloud/iaas-api-go"
	"github.com/stretchr/testify/require"
)

func TestFunctions_queryPerZone(t *testing.T) {
	serverOp := iaas.NewServerOp(testCaller)

	// prepare on is1a
	_, err := serverOp.Create(context.Background(), "is1a", &iaas.ServerCreateRequest{
		Name:     "test1",
		CPU:      1,
		MemoryMB: 1024,
	})
	require.NoError(t, err)

	// prepare on is1b
	_, err = serverOp.Create(context.Background(), "is1b", &iaas.ServerCreateRequest{
		Name:     "test1",
		CPU:      1,
		MemoryMB: 1024,
	})
	require.NoError(t, err)

	findFunc := func(ctx context.Context, zone string) ([]interface{}, error) {
		res, err := serverOp.Find(ctx, zone, nil)
		if err != nil {
			return nil, err
		}
		var results []interface{}
		for _, v := range res.Servers {
			results = append(results, v)
		}
		return results, nil
	}

	results, err := queryToZones(context.Background(), []string{"is1a", "is1b"}, findFunc)
	require.NoError(t, err)
	require.Len(t, results, 2)
}
