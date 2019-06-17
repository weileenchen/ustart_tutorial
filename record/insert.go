package record

import (
	"context"
	"strconv"

	"github.com/weileenchen/ustart_tutorial/record/recordpb"
)

func (rec *Record) Insert(ctx context.Context, req *recordpb.InsertRequest) (*recordpb.Insert, error) {
	//additional security checks (proper user format, proper email)
	intVal, err := strconv.Atoi(req.Rate)
	if err != nil {
		return nil, err
	}
	err = rec.strg.Insert(ctx, req.CarID, req.UserID, req.DateStart, intVal)
	if err != nil {
		return nil, err
	}

	return &recordpb.InsertResponse{}, nil
}
