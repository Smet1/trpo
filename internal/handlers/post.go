package handlers

import (
	"github.com/Smet1/trpo/internal/domain"
	"github.com/Smet1/trpo/internal/helpers"
	"github.com/Smet1/trpo/internal/logger"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
)

func GetCreatePostHandler(conn *sqlx.DB) http.HandlerFunc {
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

		user := &domain.User{Conn: conn}
		user.FromParsedRequest(input)

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

func GetGetPostHandler(conn *sqlx.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log := logger.GetLogger(req.Context())

		postID := chi.URLParam(req, "post_id") // from a route like /users/{post_id}
		ID, err := strconv.Atoi(postID)
		if err != nil {
			log.WithError(err).Error("can't convert post_id to string")

			helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
			return
		}

		post := &domain.Post{Conn: conn}

		posts, err := post.FindByID(int64(ID))
		if err != nil {
			log.WithError(err).Error("can't find post")

			helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
			return
		}

		helpers.Response(res, http.StatusOK, posts)
		return
	}
}
