package elasticstore

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
	"github.com/sea350/ustart_tutorial/customer/customerpb"
)

// Lookup retreives a customer doc using a certain UUID
func (estor *ElasticStore) Lookup(ctx context.Context, uuid string) (customerpb.Customer, error) {
	var customer customerpb.Customer

	termQuery := elastic.NewTermQuery("UUID", uuid)
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return customer, err
	}

	// if there are no hits, then no one exists by that uuid
	if res.Hits.TotalHits < 1 {
		return customer, errUserDoesNotExist
	}

	// if theres more than a single result then a problem has occurred
	if res.Hits.TotalHits > 1 {
		return customer, errTooManyResults
	}

	for _, elem := range res.Hits.Hits {
		//Elastic search data comes pacaged and needs to be converted into usable go structs
		data, err := elem.Source.MarshalJSON()
		if err != nil {
			return customer, err
		}

		err = json.Unmarshal(data, &customer)
		return customer, err
	}

	return customer, nil
}
