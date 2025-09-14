package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func Fail(w http.ResponseWriter, status int, msg string) {
	JSON(w, status, Error{Message: msg})
}
