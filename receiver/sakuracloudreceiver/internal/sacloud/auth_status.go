package sacloud

import (
	"context"

	"github.com/sacloud/iaas-api-go"
)

type authStatusClient interface {
	Read(context.Context) (*iaas.AuthStatus, error)
}

func getAuthStatusClient(caller iaas.APICaller) authStatusClient {
	return iaas.NewAuthStatusOp(caller)
}
