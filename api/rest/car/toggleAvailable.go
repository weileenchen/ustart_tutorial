package car

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_tutorial/car/carpb"
)

//ToggleAvailable changes availability status
func (rapi *RESTAPI) ToggleAvailable(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	avail := req.Form.Get("Available")

	lookReq := &carpb.ToggleAvailable{
		avail = lookReq
		Available: avail,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.prof.Pull(regCtx, lookReq)
	if resp != nil {
		ret["response"] = resp
	} else {
		ret["response"] = ""
	}
	if err != nil {
		ret["error"] = err.Error()
	} else {
		ret["error"] = ""
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
