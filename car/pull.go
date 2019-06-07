package car

import (
	"context"

	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

// Pull retreives all car data based off of a username
func (car *Car) Pull(ctx context.Context, req *carpb.ToggleRequest) (*carpb.ToggleResponse, error) {

	cr, err := car.strg.Lookup(ctx, req.CID)
	if err != nil {
		return nil, err
	}

	return &carpb.ToggleResponse{CarProfile: &cr}, nil

}
