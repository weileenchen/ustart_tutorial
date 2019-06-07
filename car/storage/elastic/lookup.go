package elasticstore

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

// Lookup retreives a car doc using a certain UUID
func (estor *ElasticStore) Lookup(ctx context.Context, uuid string) (carpb.Car, error) {
	var car carpb.Car

	termQuery := elastic.NewTermQuery("UUID", uuid)
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return car, err
	}

	// if there are no hits, then no one exists by that uuid
	if res.Hits.TotalHits.Value < 1 {
		return car, errCarDoesNotExist
	}

	// if theres more than a single result then a problem has occurred
	if res.Hits.TotalHits.Value > 1 {
		return car, errTooManyResults
	}

	for _, elem := range res.Hits.Hits {
		//Elastic search data comes pacaged and needs to be converted into usable go structs
		data, err := elem.Source.MarshalJSON()
		if err != nil {
			return car, err
		}

		err = json.Unmarshal(data, &car)
		return car, err
	}

	return car, nil
}
