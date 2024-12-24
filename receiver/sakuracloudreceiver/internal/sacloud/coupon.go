package sacloud

import (
	"context"
	"errors"
	"sync"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// CouponClient calls SakuraCloud coupon API
type CouponClient interface {
	Find(context.Context) ([]*iaas.Coupon, error)
}

func getCouponClient(caller iaas.APICaller) CouponClient {
	return &couponClient{caller: caller}
}

type couponClient struct {
	caller    iaas.APICaller
	accountID types.ID
	once      sync.Once
}

func (c *couponClient) Find(ctx context.Context) ([]*iaas.Coupon, error) {
	var err error
	c.once.Do(func() {
		var auth *iaas.AuthStatus

		authStatusOp := iaas.NewAuthStatusOp(c.caller)
		auth, err = authStatusOp.Read(ctx)
		if err != nil {
			return
		}
		c.accountID = auth.AccountID
	})
	if err != nil {
		return nil, err
	}
	if c.accountID.IsEmpty() {
		return nil, errors.New("getting AccountID is failed. please check your API Key settings")
	}

	couponOp := iaas.NewCouponOp(c.caller)
	searched, err := couponOp.Find(ctx, c.accountID)
	if err != nil {
		return nil, err
	}

	return searched.Coupons, nil
}
