package customer

import (
	"github.com/sea350/ustart_tutorial/customer"
)

// RESTAPI implements a REST api
// as a wrapper around the customer package.
type RESTAPI struct {
	prof *customer.Customer
}

// New creates a new customer api, given the config
func New(cfg *Config) (*RESTAPI, error) {
	prof, err := customer.New(cfg.ProfCfg)
	if err != nil {
		return nil, err
	}

	return &RESTAPI{
		prof: prof,
	}, nil
}
