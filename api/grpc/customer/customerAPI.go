package customerapi

import (
	"strconv"

	"github.com/sea350/ustart_tutorial/customer"
)

//GPRCAPI is the GRPC API implementation for customer
type GPRCAPI struct {
	prof *customer.Customer
	port string
}

// New creates a new customer api, given the config
func New(cfg *Config) (*GPRCAPI, error) {
	prof, err := customer.New(cfg.CustomerCfg)
	if err != nil {
		return nil, err
	}

	return &GPRCAPI{
		prof: prof,
		port: strconv.Itoa(cfg.Port),
	}, nil
}
