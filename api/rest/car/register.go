package car

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sea350/ustart_tutorial/car/carpb"
)

// Register wraps backend/car/register.go
func (rapi *RESTAPI) Register(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req.ParseForm()
	make := req.Form.Get("make")
	model := req.Form.Get("model")
	year := req.Form.Get("year")
	color := req.Form.Get("color")
	class := req.Form.Get("class")

	profReq := &carpb.RegisterRequest{
		Make:  make,
		Model: model,
		Year:  year,
		Color: color,
		Class: class,
	}

	ret := make(map[string]interface{})

	resp, err := rapi.car.Register(regCtx, profReq)

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
