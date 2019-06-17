package storage

import (
	elasticstore "github.com/weileenchen/ustart_tutorial/record/storage/sql"
)


func NewSQL()(config *Config) (Storage, error) {
	strg, err := &Config{SQLConfig: sqlstore.NewConfig()}
	return strg, err
}