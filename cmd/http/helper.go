package http

import (
	"encoding/json"
	"net/http"
)

func WriteWithResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	out, _ := json.Marshal(response)
	w.WriteHeader(statusCode)
	w.Write(out)
}
