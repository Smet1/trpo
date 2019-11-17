package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Smet1/trpo/internal/db"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	u := &db.User{}
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	_ = json.Unmarshal(body, u)
}
