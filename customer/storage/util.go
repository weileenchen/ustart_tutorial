package storage

import (
	"github.com/google/uuid"
	"github.com/sea350/ustart_tutorial/customer/customerpb"
)

// Dummy is a dummy implementation of Register, managed in-memory.
// Not to be used in Production.
type Dummy struct {
	customers  map[string]*customerpb.Customer // maps UUID to a user object
	emailToID map[string]string             // maps email to uuid
}

// GetPassword retrieves a password, given a uuid
func GetPassword(userUUID uuid.UUID) {

}
