package customer

import (
	"context"

	"github.com/sea350/ustart_tutorial/customer/customerpb"
)

// Pull retreives all customer data based off of a username
func (customer *Customer) Pull(ctx context.Context, req *customerpb.PullRequest) (*customerpb.PullResponse, error) {

	cust, err := customer.strg.Lookup(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	return &customerpb.PullResponse{CustomerProfile: &cust}, nil

}
