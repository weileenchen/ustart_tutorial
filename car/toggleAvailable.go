package car

import (
	"context"

	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

// ToggleAvailable changes status from true/false to false/true depending on initial response
func (car *Car) ToggleAvailable(ctx context.Context, req carpb.ToggleRequest) (carpb.ToggleResponse, error) {
	//start diverging in the storage
	//check current status -- storage function -- check car availability
	//then have another function which updates availability
	//interface layer mirror those changes
	//toggle available will use those extra functions to solve task
	response, err := car.strg.CheckAvailabilityStatus(ctx, req.CID)
	var resp carpb.ToggleResponse
	if err != nil {
		return resp, err
	}
	resp.NewStatus, err = car.strg.UpdateAvailabilityStatus(ctx, req.CID, response)
	return resp, err
}
