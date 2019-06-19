package customer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_tutorial/customer/customerpb"
)

// Register wraps backend/customer/register.go
func (rapi *RESTAPI) Register(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	fname := req.Form.Get("firstname")
	lname := req.Form.Get("lastname")
	dob := req.Form.Get("dob")

	profReq := &customerpb.RegisterRequest{
		FirstName: fname,
		LastName:  lname,
		DOB:       dob,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.prof.Register(regCtx, profReq)
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

	data, err := json.Marshal(ret) //takes information and converts to json
	if err != nil {
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
