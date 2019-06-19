package record

import "github.com/weileenchen/ustart_tutorial/record/storage"

// Config determines the runtime behavior of the SQL-backed customer server
type Config struct {
	useDummy      bool
	StorageConfig *storage.Config
}
