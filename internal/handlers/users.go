package handlers

import (
	"encoding/json"
	"github.com/Smet1/trpo/internal/logger"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"

	"github.com/Smet1/trpo/internal/db"
)

func GetCreateUserHandler(conn *sqlx.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log := logger.GetLogger(req.Context())
		u := &db.User{}
		body, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		_ = json.Unmarshal(body, u)

		err := u.Insert(conn)
		if err != nil {
			log.WithError(err).Error("can't create user")
		}
	}
}
