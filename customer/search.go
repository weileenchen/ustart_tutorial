package customer

import (
	"context"

	"github.com/sea350/ustart_tutorial/customer/customerpb"
)

// Search retreives a list of minimal customer data based off search queries
func (customer *Customer) Search(ctx context.Context, req *customerpb.SearchRequest) (*customerpb.SearchResponse, error) {

	queryArr := []string{req.FirstName, req.LastName, req.DOB}
	// filterArr := strings.FieldsFunc(req.Filters, split)

	//TODO:
	//enable filters
	//set up a real must map
	//add comments

	searchName := true

	ids, err := customer.strg.Search(ctx, queryArr, searchName, make(map[string][]string), req.Scroll)
	if err != nil {
		return nil, err
	}

	var res []*customerpb.Customer
	var resErr error
	for _, id := range ids {
		cust, err := customer.strg.Lookup(ctx, id)
		if err != nil {
			res = append(res, &cust)
		} else {
			resErr = ErrProblemLoadingcustomer
		}
	}
	return &customerpb.SearchResponse{Results: res}, resErr
}
