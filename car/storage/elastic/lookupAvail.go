package elasticstore

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

// LookupAvail retreives a car availability status doc using a carID
func (estor *ElasticStore) LookupAvail(ctx context.Context, CarID string) (bool, error) {
	var car carpb.Car

	termQuery := elastic.NewTermQuery("CarID", CarID)
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return car.Available, err
	}

	//keep in mind empty car.Available
	// if there are no hits, then no one exists by that uuid
	if res.Hits.TotalHits.Value < 1 {
		return car.Available, errCarDoesNotExist
	}

	// if theres more than a single result then a problem has occurred
	if res.Hits.TotalHits.Value > 1 {
		return car.Available, errTooManyResults
	}

	for _, elem := range res.Hits.Hits {
		//Elastic search data comes pacaged and needs to be converted into usable go structs
		data, err := elem.Source.MarshalJSON()
		if err != nil {
			return car.Available, err
		}

		err = json.Unmarshal(data, &car)
		return car.Available, err
	}

	return car.Available, nil
}
