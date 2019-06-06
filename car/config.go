package car

import "github.com/weileenchen/ustart_tutorial/car/storage"

// Config determines the runtime behavior of the Elastic-backed car server
type Config struct {
	useDummy      bool
	StorageConfig *storage.Config
}
