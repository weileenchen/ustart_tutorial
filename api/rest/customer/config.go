package customer

import (
	"github.com/sea350/ustart_tutorial/customer"
)

// Config for customer api
type Config struct {
	ProfCfg *customer.Config
	Port    int //Recomended 5002
}
