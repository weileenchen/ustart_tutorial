package car

import "errors"

var (
	//errCarExists This user already has a car
	errCarExists = errors.New("This user already has a car")

	//errProblemLoadingCar This user already has a car
	errProblemLoadingCar = errors.New("There was a problem loading one or more cars in this list")
)
