package recordapi

import (
	"strconv"

	"github.com/weileenchen/ustart_tutorial/record"
)

//GPRCAPI is the GRPC API implementation for record
type GPRCAPI struct {
	record *record.Record
	port   string
}

// New creates a new customer api, given the config
func New(cfg *Config) (*GPRCAPI, error) {
	rec, err := record.New(cfg.RecordCfg)
	if err != nil {
		return nil, err
	}

	return &GPRCAPI{
		record: prof,
		port:   strconv.Itoa(cfg.Port),
	}, nil
}
