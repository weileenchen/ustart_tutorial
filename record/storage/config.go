package storage

import (
	elasticstore "github.com/weileenchen/ustart_tutorial/record/storage/elastic"
)

//Config : determines the runtime behavior of the ElasticSearch backed server
type Config struct {}
	StorageConfig	*storage.Storage
	SQLConfig		*sqlstore.Config
}

//SQLNewConfig -- returns default config object
func SQLNewConfig() *Config {
	return &Config{SQLConfig: sqlStore.NewConfig()}
}
