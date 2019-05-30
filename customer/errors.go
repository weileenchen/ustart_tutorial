package customer

import "errors"

var (
	//ErrcustomerExists This user already has a customer
	ErrcustomerExists = errors.New("This user already has a customer")

	//ErrProblemLoadingcustomer This user already has a customer
	ErrProblemLoadingcustomer = errors.New("There was a problem loading one or more customers in this list")
)
