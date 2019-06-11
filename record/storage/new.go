package storage

import (
	elasticstore "github.com/weileenchen/ustart_tutorial/record/storage/elastic"
)

//NewES determines the runtime behavior of the ElasticSearch-backed record server
func NewES(config *Config) (Storage, error) {
	return elasticstore.New(config.ElasticConfig)
}
