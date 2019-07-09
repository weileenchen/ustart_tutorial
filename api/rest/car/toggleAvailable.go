package car

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

//ToggleAvailable changes availability status
func (rapi *CARRESTAPI) ToggleAvailable(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()

	carID := req.Form.Get("carID")
	var togReq carpb.ToggleRequest
	togReq.CID = carID

	ret := make(map[string]interface{})

	resp, err := rapi.car.ToggleAvailable(regCtx, togReq)

	if resp.NewStatus != nil {
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
