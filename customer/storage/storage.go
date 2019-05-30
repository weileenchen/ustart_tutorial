package storage

import (
	"context"

	"github.com/sea350/ustart_tutorial/customer/customerpb"
)

// Storage is a database-agnostic interface for persisting customer data
type Storage interface {
	Register(context.Context, string, string, string, string) error
	Lookup(context.Context, string) (customerpb.Customer, error)
	Search(context.Context, []string, bool, map[string][]string, string) ([]string, error)
	// rest of the functions
	ErrUserDoesNotExist() error
	ErrTooManyResults() error
}
