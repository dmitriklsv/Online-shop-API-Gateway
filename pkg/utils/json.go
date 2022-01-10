package utils

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, status int64, postBody any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	json.NewEncoder(w).Encode(postBody)
}
