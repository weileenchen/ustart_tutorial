package car

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_tutorial/car/carpb"
)

//Search 
func (rapi *RESTAPI) Search(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	id := req.Form.Get("CarID")

	lookReq := &carpb.Search(id){
		CarID: id,
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
