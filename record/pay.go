package record

import "context"

//Pay -- Request looks for most recent unpaid history entry; pays it off with that quantity
func (record *Record) Pay(context.Context, []string, bool, map[string][]string, string) ([]string, error) {
	...
}
