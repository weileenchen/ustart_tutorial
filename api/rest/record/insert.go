package record

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/weileenchen/ustart_tutorial/record/recordpb"
)

// Pull wraps backend/customer/pull.go
func (rapi *RESTAPI) Pull(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//NOTE: this method of retrieving data from a REST request should only be used for GET requests
	//later on you will be shown the difference between what is and when you should use GET or POST
	req.ParseForm()
	carID := req.Form.Get("carID")
	uID := req.Form.Get("UserID")
	date := req.Form.Get("startDate")
	rate := req.Form.Get("rate")

	insertReq := &recordpb.InsertRequest{
		CarID:     carID,
		UserID:    uID,
		DateStart: date,
		Rate:      rate,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.record.Pull(regCtx, lookReq)
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
		log.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
