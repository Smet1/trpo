package handlers

import (
	"github.com/Smet1/trpo/internal/domain"
	"github.com/Smet1/trpo/internal/helpers"
	"net/http"

	"github.com/Smet1/trpo/internal/logger"
	"github.com/jmoiron/sqlx"
)

func GetCreateUserHandler(conn *sqlx.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log := logger.GetLogger(req.Context())
		input := &helpers.User{}
		err := input.ParseFromRequest(req.Body)
		if err != nil {
			log.WithError(err).Error("wrong request body")

			helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
			return
		}
		defer req.Body.Close()

		user := &domain.User{}
		user.FromParsedRequest(input, conn)

		err = user.Validate()
		if err != nil {
			log.WithError(err).Error("not valid data")

			helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
			return
		}

		err = user.Create()
		if err != nil {
			log.WithError(err).Error("can't create user")

			helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
			return
		}

		log.WithField("user", user).Info("user created")

		helpers.Response(res, http.StatusCreated, user.ToResponse())
		return
	}
}
