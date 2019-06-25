package car

import (
	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

// ToggleAvailable changes status from true/false to false/true depending on initial response
func (car *Car) ToggleAvailable(req carpb.ToggleRequest) bool {
	//start diverging in the storage
	//check current status -- storage function -- check car availability
	//then have another function which updates availability
	//interface layer mirror those changes
	//toggle available will use those extra functions to solve task
	bool response = CheckAvailabilityStatus(req)

	if response == true {
		UpdateAvailabilityStatus(false)
		return false
	} else {
		UpdateAvailabilityStatus(true)
		return true
	}
}
