package customer

import (
	"context"

	"github.com/sea350/ustart_tutorial/customer/customerpb"
)

// Register is a generic register function that registers a user in a database
func (customer *Customer) Register(ctx context.Context, req *customerpb.RegisterRequest) (*customerpb.RegisterResponse, error) {

	uuid := randString(32)

	_, err := customer.strg.Lookup(ctx, uuid)
	if err != nil && err != customer.strg.ErrUserDoesNotExist() {
		return nil, err
	}
	if err == nil {
		return nil, errCustomerExists
	}

	err = customer.strg.Register(ctx, uuid, req.FirstName, req.LastName, req.DOB)
	if err != nil {
		return nil, err
	}

	return &customerpb.RegisterResponse{}, nil

}
