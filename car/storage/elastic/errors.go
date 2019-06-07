package elasticstore

import "errors"

//Having errors globalized like this makes it easier for higher level packages to identify specific problems
//and implement logic for special error cases
//the variables are not exported so they cant be modified by anything outside the package

var (
	// errCarDoesNotExist user doesnt exist
	errCarDoesNotExist = errors.New("Car does not exist")

	// errTooManyResults if more than one result per email shows
	errTooManyResults = errors.New("Too many results, a crititcal error has occurred")
)

//ErrCarDoesNotExist returns a standardized error
func (estor *ElasticStore) ErrCarDoesNotExist() error {
	return errCarDoesNotExist
}

//ErrTooManyResults returns a standardized error
func (estor *ElasticStore) ErrTooManyResults() error {
	return errTooManyResults
}
