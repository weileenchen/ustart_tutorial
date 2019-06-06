package storage

// Storage is a database-agnostic interface for persisting customer data
// import (
// 	"context"

// 	"github.com/sea350/ustart_tutorial/record/recordpb"
// )

type Storage interface {
	ErrUserDoesNotExist() error
	ErrTooManyResults() error
}
