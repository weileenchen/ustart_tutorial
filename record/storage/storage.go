package storage

import (
	"context"

	"github.com/sea350/ustart_tutorial/record/recordpb"
)

//Storage : interface for data
type Storage interface {
	//HistoryRequest pulls users rental history
	History(context.Context, string) (recordpb.Record, error)
	
	//PayRequest looks for most recent unpaid history entry; pays it off with that quantity
	//fix later, figure out how to implement Pay
	Pay(context.Context, []string, bool, map[string][]string, string) ([]string, error)

	//NewRequest done in SQL -- submit a new row with the data
	Insert(ctx context.Context, string, string, string, int) error

	ErrRecordDoesNotExist() error
	ErrTooManyResults() error
}
