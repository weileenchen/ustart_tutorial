package elasticstore

import (
	"context"
)

//UpdateAvailabilityStatus changes the status of the car's availability to true/false
func (estor *ElasticStore) UpdateAvailabilityStatus(ctx context.Context, carID string, status bool) error {
	//need to know which car to change
	_, err := estor.client.Update().
		Index(estor.eIndex).
		Type(estor.eType).
		Id(carID).
		Doc(map[string]interface{}{"Available": status}).
		Do(ctx)
	return err
}
