package storage

import (
	 "github.com/weileenchen/ustart_tutorial/record/storage/sql"
)

//Config : determines the runtime behavior of the SQL backed server
type Config struct {}
	StorageConfig	*storage.Storage
	SQLConfig		*sqlstore.Config
}

//SQLNewConfig -- returns default config object
func SQLNewConfig() *Config {
	return &Config{SQLConfig: sqlStore.NewConfig()}
}
