package elasticstore

import (
	"context"

	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

//Register creates a new ES document for a new registering car
func (estor *ElasticStore) Register(ctx context.Context, make string, model string, year string, color string, class string) error {

	//Lock just to make sure no two people can sign up with the same username at the same time
	newUserLock.Lock()
	defer newUserLock.Unlock()

	_, err := estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(carpb.Car{
			Make:  make,
			Model: model,
			Year:  year,
			Color: color,
			Class: class,
		}).
		//Double check this id
		Id(make).
		Do(ctx)

	return err
}
