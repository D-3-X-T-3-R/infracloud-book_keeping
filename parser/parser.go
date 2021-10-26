package parser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseJsonBody(req *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(req.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
