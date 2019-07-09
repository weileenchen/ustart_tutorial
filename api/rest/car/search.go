package car

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/weileenchen/ustart_tutorial/car/carpb"
)

//Search implements search
func (rapi *CARRESTAPI) Search(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	firstName := req.Form.Get("FirstName")
	lastName := req.Form.Get("LastName")
	dob := req.Form.Get("DOB")
	scroll := req.Form.Get("Scroll")

	lookReq := &carpb.Search(firstName, lastName, dob, scroll)

	ret := make(map[string]interface{})

	resp, err := rapi.car.Search(regCtx, lookReq)
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
