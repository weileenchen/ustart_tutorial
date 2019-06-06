package car

import "errors"

var (

	//errProblemLoadingCustomer This user already has a car
	errProblemRetrievingCar = errors.New("There was a problem loading one or more cars in this list")
)
