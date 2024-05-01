package response

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Json(w http.ResponseWriter, data Data) error {
	c, e := json.Marshal(data)
	if e != nil {
		return e
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(data.Code)
	w.Write(c)
	return nil
}
