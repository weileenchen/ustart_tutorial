package elasticstore

import (
	"context"

	"github.com/olivere/elastic"
)

// ElasticStore implements the storage interface for the customer module
type ElasticStore struct {
	client *elastic.Client
	eIndex string
	eType  string
}

// New returns a new Eclient elasticstore service
func New(cfg *Config) (*ElasticStore, error) {

	//generating elastic client based on config infrmation
	client, err := elastic.NewClient(elastic.SetURL(cfg.ElasticAddr))
	if err != nil {
		panic(err)
	}

	//setting up our storage interface
	ecl := &ElasticStore{
		client: client,
		eIndex: cfg.EIndex,
		eType:  cfg.EType,
	}

	//NOTE: when using a non-imported context you generally want to use a time out
	//this was not done in this case
	//observe the difference between pingCtx and the implented default context
	// pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	// defer cancel()
	_, _, err = ecl.client.Ping(cfg.ElasticAddr).Do(context.Background())
	if err != nil {
		panic(err)
	}

	//Make sure the necessary Indexes exist in your ES node
	exists, err := ecl.client.IndexExists(ecl.eIndex).Do(context.Background())
	if err != nil {
		panic(err)
	}
	//If they don't exist not create them
	//doing this is the equivalent of `CREATE TABLE IF NOT EXISTS` in sql
	if !exists {
		createIndex, err := ecl.client.CreateIndex(ecl.eIndex).BodyString(mapping(cfg.EIndex)).Do(context.Background()) //DONT FORGET TO ADD MAPPTING LATER
		if err != nil {
			panic(err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
			panic(err)
		}
	}

	return ecl, err
}
