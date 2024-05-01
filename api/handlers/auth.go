package handlers

import (
	"net/http"
	response "qeymat/api/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {

	response.Json(w, response.Data{
		Code:    200,
		Message: "Was successful.",
	})
}
