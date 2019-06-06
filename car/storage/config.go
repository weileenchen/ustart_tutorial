package storage

import (
	elasticstore "github.com/weileenchen/ustart_tutorial/car/storage/elastic"
)

// Config determines the runtime behavior of the an either SQL or ElasticSearch backed car server
type Config struct {
	useDummy      bool
	ElasticConfig *elasticstore.Config
	// SQLConfig     *sqlstore.Config
}

// ESNewConfig returns a default config object
func ESNewConfig() *Config {
	return &Config{ElasticConfig: elasticstore.NewConfig()}
}
