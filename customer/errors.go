package customer

import "errors"

var (
	//errCustomerExists This user already has a customer
	errCustomerExists = errors.New("This user already has a customer")

	//errProblemLoadingCustomer This user already has a customer
	errProblemLoadingCustomer = errors.New("There was a problem loading one or more customers in this list")
)
