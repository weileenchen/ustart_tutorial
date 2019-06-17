package recordapi

import (
	"strconv"

	"github.com/weileenchen/ustart_tutorial/record"
)

//RESTAPI is the REST API implementation for record
type RESTAPI struct {
	record *record.Record
	port   string
}

// New creates a new customer api, given the config
func New(cfg *Config) (*RESTAPI, error) {
	rec, err := record.New(cfg.RecordCfg)
	if err != nil {
		return nil, err
	}

	return &RESTAPI{
		record: prof,
		port:   strconv.Itoa(cfg.Port),
	}, nil
}
