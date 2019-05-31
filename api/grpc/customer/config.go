package customerapi

import (
	"github.com/sea350/ustart_tutorial/customer"
)

// Config for auth api
type Config struct {
	CustomerCfg *customer.Config
	Port       int //Recomended 5101
}
