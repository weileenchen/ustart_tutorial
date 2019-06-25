package car

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//ToggleAvailable changes availability status
func (rapi *RESTAPI) ToggleAvailable(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()

	carID := req.Form.Get("carID")

	ret := make(map[string]interface{})

	resp, err := rapi.car.ToggleAvailable(regCtx, carID)

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
