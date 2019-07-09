package car

import (
	"github.com/weileenchen/ustart_tutorial/car"
)

// CARRESTAPI implements a REST api as a wrapper around the customer package.
type CARRESTAPI struct {
	car *car.Car
}

// New creates a new car api, given the config
func New(cfg *Config) (*CARRESTAPI, error) {
	car, err := car.New(cfg.CarCfg)
	if err != nil {
		return nil, err
	}

	return &CARRESTAPI{
		car: car,
	}, nil
}
