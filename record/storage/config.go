package storage

import (
	elasticstore "github.com/weileenchen/ustart_tutorial/record/storage/elastic"
)

//Config : determines the runtime behavior of the ElasticSearch backed server
type Config struct {
	useDummy      bool
	ElasticConfig *elasticstore.Config
}

//ESNewConfig -- returns default config object
func ESNewConfig() *Config {
	return &Config{ElasticConfig: elasticstore.NewConfig()}
}
