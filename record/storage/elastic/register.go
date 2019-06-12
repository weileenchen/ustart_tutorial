package elasticstore

import (
	"context"

	"github.com/weileenchen/ustart_tutorial/record/recordpb"
)

//Register creates a new ES document for a new record
func (estor *ElasticStore) Register(ctx context.Context, carID string, userID string, dateStart string, rate string) error {

	//Lock just to make sure no two people can sign up with the same userID at the same time
	newUserLock.Lock()
	defer newUserLock.Unlock()

	_, err := estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(recordpb.Record{
			CarID:     carID,
			UserID:    userID,
			DateStart: dateStart,
			//converts rate from string to float64 through the logic layer
			//Rate:      rate,
		}).
		Id(userID).
		Do(ctx)

	return err
}
