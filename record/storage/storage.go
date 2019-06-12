package storage

import (
	"context"

	"github.com/sea350/ustart_tutorial/record/recordpb"
)

//Storage : interface for data
type Storage interface {
	New(context.Context, string, string, string, string) error
	History(context.Context, string) (recordpb.Record, error)
	//fix later, figure out how to implement Pay
	Pay(context.Context, []string, bool, map[string][]string, string) ([]string, error)

	ErrRecordDoesNotExist() error
	ErrTooManyResults() error
}
