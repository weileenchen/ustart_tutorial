package elasticstore

import "context"

//CheckAvailabilityStatus finds and returns the availability status of the car
func (estor *ElasticStore) CheckAvailabilityStatus(ctx context.Context, request string) (bool, error) {
	return estor.Lookup(ctx, request)
}
