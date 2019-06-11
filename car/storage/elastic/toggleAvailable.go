package elasticstore

import (
	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

// ToggleAvailable changes status from true/false to false/true depending on initial response
func (estor *ElasticStore) ToggleAvailable(carResponse carpb.ToggleResponse) {
	if carResponse.NewStatus == true {
		carResponse.NewStatus = false
	} else {
		carResponse.NewStatus = true
	}
}
