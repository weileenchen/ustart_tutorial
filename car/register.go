package car

import (
	"context"

	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

// Register is a generic register function that registers a user in a database
func (car *Car) Register(ctx context.Context, req *carpb.RegisterRequest) (*carpb.RegisterResponse, error) {

	CID := randString(32)

	_, err := car.strg.Lookup(ctx, CID)
	if err != nil && err != car.strg.ErrCarDoesNotExist() {
		return nil, err
	}
	if err == nil {
		return nil, errCarExists
	}

	err = car.strg.Register(ctx, CID, req.Make, req.Model, req.Year, req.Color, req.Class)
	if err != nil {
		return nil, err
	}

	return &carpb.RegisterResponse{}, nil

}
