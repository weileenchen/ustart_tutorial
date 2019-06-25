package customer

import (
	"github.com/sea350/ustart_tutorial/customer"
)

// RESTAPI implements a REST api
// as a wrapper around the customer package.
type RESTAPI struct {
	cust *customer.Customer
}

// New creates a new customer api, given the config
func New(cfg *Config) (*RESTAPI, error) {
	cust, err := customer.New(cfg.ProfCfg)
	if err != nil {
		return nil, err
	}

	return &RESTAPI{
		cust: cust,
	}, nil
}
