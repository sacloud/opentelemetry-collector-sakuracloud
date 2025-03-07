package sacloud

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
)

// BillClient calls SakuraCloud bill API
type BillClient interface {
	Read(context.Context) (*iaas.Bill, error)
}

func getBillClient(caller iaas.APICaller) BillClient {
	return &billClient{caller: caller}
}

type billClient struct {
	caller    iaas.APICaller
	accountID types.ID
	once      sync.Once
}

func (c *billClient) Read(ctx context.Context) (*iaas.Bill, error) {
	var err error
	c.once.Do(func() {
		var auth *iaas.AuthStatus

		authStatusOp := iaas.NewAuthStatusOp(c.caller)
		auth, err = authStatusOp.Read(ctx)
		if err != nil {
			return
		}
		if !auth.ExternalPermission.PermittedBill() {
			err = fmt.Errorf("account doesn't have permissions to use the Billing API")
		}
		c.accountID = auth.AccountID
	})
	if err != nil {
		return nil, err
	}
	if c.accountID.IsEmpty() {
		return nil, errors.New("getting AccountID is failed. please check your API Key settings")
	}

	billOp := iaas.NewBillOp(c.caller)
	searched, err := billOp.ByContract(ctx, c.accountID)
	if err != nil {
		return nil, err
	}

	var bill *iaas.Bill
	for i := range searched.Bills {
		b := searched.Bills[i]
		if i == 0 || bill.Date.Before(b.Date) {
			bill = b
		}
	}
	return bill, nil
}
