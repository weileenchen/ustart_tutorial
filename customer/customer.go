package customer

import (
	"github.com/sea350/ustart_tutorial/customer/storage"
)

//Customer is an implementation of the customer service as defined in service.proto
type Customer struct {
	strg          storage.Storage
	defaultAvatar string
	defaultBanner string
}

// New returns a new Eclient customer service
func New(cfg *Config) (*Customer, error) {
	// if cfg.useDummy

	storg, err := storage.NewES(cfg.StorageConfig)

	cust := &Customer{
		strg: storg,
	}

	return cust, err
}
