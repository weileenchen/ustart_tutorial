package record

import (
	"context"

	"github.com/weileenchen/ustart_tutorial/record/recordpb"
)

//History -- request pulls users rental history
func (record *Record) History(context.Context, string) (recordpb.Record, error) {

}
