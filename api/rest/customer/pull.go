package customer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sea350/ustart_tutorial/customer/customerpb"
)

// Pull wraps backend/customer/pull.go
func (rapi *RESTAPI) Pull(w http.ResponseWriter, req *http.Request) {
	//old method, should be revised
	//don't designate our own context
	// regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()

	//NOTE: this method of retrieving data from a REST request should only be used for GET requests
	//later on you will be shown the difference between what is and when you should use GET or POST
	req.ParseForm()
	uuid := req.Form.Get("uuid")

	lookReq := &customerpb.PullRequest{
		UUID: uuid,
	}

	//value interface means it's agnostic
	ret := make(map[string]interface{})

	//take info, pack it up and send it out
	//use req.Context instead of designating our own context
	resp, err := rapi.cust.Pull(req.Context(), lookReq)
	if resp != nil {
		ret["response"] = resp
	} else {
		//cannot package a nil response, so instead package a blank string
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

	//Fprint -- outputs the same way but specifies which input/output socket to dump everything into
	fmt.Fprintln(w, string(data))
}
