package elasticstore

import "errors"

//Having errors globalized like this makes it easier for higher level packages to identify specific problems
//and implement logic for special error cases
//the variables are not exported so they cant be modified by anything outside the package

var (
	// errUserDoesNotExist user doesnt exist
	errUserDoesNotExist = errors.New("User does not exist")

	// errTooManyResults if more than one result per email shows
	errTooManyResults = errors.New("Too many results, a crititcal error has occurred")
)

//ErrUserDoesNotExist returns a standardized error
func (estor *ElasticStore) ErrUserDoesNotExist() error {
	return errUserDoesNotExist
}

//ErrTooManyResults returns a standardized error
func (estor *ElasticStore) ErrTooManyResults() error {
	return errTooManyResults
}
