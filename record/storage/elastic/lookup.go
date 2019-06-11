package elasticstore

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic"
	"github.com/weileenchen/ustart_tutorial/record/recordpb"
)

// Lookup retreives a record doc using a certain UUID
func (estor *ElasticStore) Lookup(ctx context.Context, uuid string) (recordpb.Record, error) {
	var record recordpb.Record

	termQuery := elastic.NewTermQuery("UUID", uuid)
	res, err := estor.client.Search().
		Index(estor.eIndex).
		Query(termQuery).
		Do(ctx)

	if err != nil {
		return record, err
	}

	// if there are no hits, then no one exists by that uuid
	if res.Hits.TotalHits.Value < 1 {
		return record, errRecordDoesNotExist
	}

	// if theres more than a single result then a problem has occurred
	if res.Hits.TotalHits.Value > 1 {
		return record, errTooManyResults
	}

	for _, elem := range res.Hits.Hits {
		//Elastic search data comes pacaged and needs to be converted into usable go structs
		data, err := elem.Source.MarshalJSON()
		if err != nil {
			return record, err
		}

		err = json.Unmarshal(data, &record)
		return record, err
	}

	return record, nil
}
