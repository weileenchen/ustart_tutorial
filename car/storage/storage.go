package storage

import (
	"context"

	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

// Storage is a database-agnostic interface for persisting car data
type Storage interface {
	Register(context.Context, string, string, string, string, string, string) error
	Search(context.Context, []string, bool, map[string][]string, string) ([]string, error)
	Lookup(context.Context, string) (carpb.Car, error)
	//ToggleAvailable(string) bool
	// rest of the functions
	ErrCarDoesNotExist() error
	ErrTooManyResults() error
}
