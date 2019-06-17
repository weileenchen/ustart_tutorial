package record

import (
	"github.com/weileenchen/ustart_tutorial/record/storage"
)

//Record is an implementation of the record service as defined in service.proto
type Record struct {
	strg          storage.Storage
	defaultAvatar string
	defaultBanner string
}

// New returns a new Eclient record service
func New(cfg *Config) (*Record, error) {

	return nil, nil
}
