package storage

import (
	elasticstore "github.com/sea350/ustart_tutorial/customer/storage/elastic"
)

// NewES determines the runtime behavior of the ElasticSearch-backed customer server
func NewES(config *Config) (Storage, error) {
	return elasticstore.New(config.ElasticConfig)
}
