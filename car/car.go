package car

import (
	"github.com/weileenchen/ustart_tutorial/car/storage"
)

//Car is an implementation of the car service as defined in car.proto
type Car struct {
	strg          storage.Storage
	defaultAvatar string
	defaultBanner string
	//should I add more string vars?
}

// New returns a new Eclient car service
func New(cfg *Config) (*Car, error) {
	// if cfg.useDummy

	storg, err := storage.NewES(cfg.StorageConfig)

	car := &Car{
		strg: storg,
	}

	return car, err
}
