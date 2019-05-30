package customer

import "github.com/sea350/ustart_tutorial/customer/storage"

// Config determines the runtime behavior of the Elastic-backed customer server
type Config struct {
	useDummy      bool
	StorageConfig *storage.Config
}
