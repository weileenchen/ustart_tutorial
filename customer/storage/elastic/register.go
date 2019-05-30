package elasticstore

import (
	"context"

	"github.com/sea350/ustart_tutorial/customer/customerpb"
)

//Register creates a new ES document for a new registering customer
func (estor *ElasticStore) Register(ctx context.Context, uuid string, firstname string, lastname string, dob string) error {

	//Lock just to make sure no two people can sign up with the same username at the same time
	newUserLock.Lock()
	defer newUserLock.Unlock()

	_, err := estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(customerpb.Customer{
			UUID:      uuid,
			FirstName: firstname,
			LastName:  lastname,
			DOB:       dob,
		}).
		Id(uuid).
		Do(ctx)

	return err
}
