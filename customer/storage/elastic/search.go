package elasticstore

import (
	"context"
	"io"

	"github.com/olivere/elastic"
)

// Search searches for... something... i think
func (estor *ElasticStore) Search(ctx context.Context, searchTerms []string, searchName bool, mustMap map[string][]string, scrollID string) ([]string, error) {
	query := elastic.NewBoolQuery()

	for _, term := range searchTerms {
		if searchName {
			query = query.Should(elastic.NewFuzzyQuery("FirstName", term).Fuzziness("AUTO"))
			query = query.Should(elastic.NewFuzzyQuery("LastName", term).Fuzziness("AUTO"))
		}
	}

	results := []string{}

	scroll := estor.client.Scroll().
		Index(estor.eIndex).
		Query(query).
		Size(10).
		Sort("_score", false)

	if scrollID != "" {
		scroll = scroll.ScrollId(scrollID)
	}

	res, err := scroll.Do(ctx)
	if (err != nil && err != io.EOF) || res == nil {
		return results, err
	}

	for _, element := range res.Hits.Hits {
		results = append(results, element.Id)
	}
	return results, nil
}
