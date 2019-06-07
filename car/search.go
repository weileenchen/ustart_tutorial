package car

import (
	"context"

	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

// Search retreives a list of minimal customer data based off search queries
func (caer *Car) Search(ctx context.Context, req *carpb.SearchRequest) (*carpb.SearchResponse, error) {

	queryArr := []string{req.FirstName, req.LastName, req.DOB}
	// filterArr := strings.FieldsFunc(req.Filters, split)

	//TODO:
	//enable filters
	//set up a real must map
	//add comments

	searchName := true

	ids, err := car.strg.Search(ctx, queryArr, searchName, make(map[string][]string), req.Scroll)
	if err != nil {
		return nil, err
	}

	var res []*carpb.Car
	var resErr error
	for _, id := range ids {
		cust, err := car.strg.Lookup(ctx, id)
		if err != nil {
			res = append(res, &cust)
		} else {
			resErr = errProblemLoadingCar
		}
	}
	return &carpb.SearchResponse{Results: res}, resErr
}
