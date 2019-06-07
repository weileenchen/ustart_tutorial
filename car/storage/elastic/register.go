package elasticstore

import (
	"context"

	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

//Register creates a new ES document for a new registering car
func (estor *ElasticStore) Register(ctx context.Context, cid string, make string, model string, year string, color string, class string) error {

	//Lock just to make sure no two people can sign up with the same username at the same time
	newUserLock.Lock()
	defer newUserLock.Unlock()

	_, err := estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(carpb.Car{
			CID:   cid,
			Make:  make,
			Model: model,
			Year:  year,
			Color: color,
			Class: class,
		}).
		Id(cid).
		Do(ctx)

	return err
}
