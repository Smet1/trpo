package handlers

import (
	"net/http"

	"github.com/Smet1/trpo/internal/domain"
	"github.com/Smet1/trpo/internal/helpers"
	"github.com/Smet1/trpo/internal/logger"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Conn *sqlx.DB
}

func (uh *User) CreateUser(res http.ResponseWriter, req *http.Request) {
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
	user.FromParsedRequest(input)

	err = user.Validate()
	if err != nil {
		log.WithError(err).Error("not valid data")

		helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
		return
	}

	err = user.Create(uh.Conn)
	if err != nil {
		log.WithError(err).Error("can't create user")

		helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
		return
	}

	log.WithField("user", user).Info("user created")

	helpers.Response(res, http.StatusCreated, user.ToResponse())
	return
}

func (uh *User) GetUser(res http.ResponseWriter, req *http.Request) {
	log := logger.GetLogger(req.Context())

	username := chi.URLParam(req, "username") // from a route like /users/{username}
	user := &domain.User{}

	err := user.GetByUsername(username, uh.Conn)
	if err != nil {
		log.WithError(err).Error("can't find user")

		helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
		return
	}

	helpers.Response(res, http.StatusOK, user.ToResponse())
	return
}

func (uh *User) Auth(res http.ResponseWriter, req *http.Request) {
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

	_, err = user.Auth(input.Login, input.Password, uh.Conn)
	if err != nil {
		helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: "wrong password or login"})
		return
	}

	helpers.Response(res, http.StatusOK, user.ToResponse())
	return
}
