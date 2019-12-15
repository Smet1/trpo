package helpers

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func Response(res http.ResponseWriter, status int, body interface{}) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(status)
	binary, err := json.Marshal(body)
	if err != nil {
		return
	}

	_, err = res.Write(binary)
	if err != nil {
		return
	}
}
