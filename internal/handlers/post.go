package handlers

import (
	"net/http"
	"strconv"

	"github.com/Smet1/trpo/internal/domain"
	"github.com/Smet1/trpo/internal/helpers"
	"github.com/Smet1/trpo/internal/logger"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

type Posts struct {
	Conn *sqlx.DB
}

func (ph *Posts) CreatePost(res http.ResponseWriter, req *http.Request) {
	log := logger.GetLogger(req.Context())
	input := &helpers.Post{}
	err := input.ParseFromRequest(req.Body)
	if err != nil {
		log.WithError(err).Error("wrong request body")

		helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
		return
	}
	defer req.Body.Close()

	user := &domain.Post{Conn: ph.Conn}
	user.FromParsedRequest(input)

	err = user.Create()
	if err != nil {
		log.WithError(err).Error("can't create post")

		helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
		return
	}

	log.WithField("post", user).Info("post created")

	helpers.Response(res, http.StatusCreated, user.ToResponse())
	return
}

func (ph *Posts) GetPost(res http.ResponseWriter, req *http.Request) {
	log := logger.GetLogger(req.Context())

	postID := chi.URLParam(req, "post_id") // from a route like /users/{post_id}
	ID, err := strconv.Atoi(postID)
	if err != nil {
		log.WithError(err).Error("can't convert post_id to string")

		helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
		return
	}

	post := &domain.Post{Conn: ph.Conn}

	posts, err := post.FindByID(int64(ID))
	if err != nil {
		log.WithError(err).Error("can't find post")

		helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
		return
	}

	helpers.Response(res, http.StatusOK, posts)
	return
}

func (ph *Posts) GetPosts(res http.ResponseWriter, req *http.Request) {
	log := logger.GetLogger(req.Context())

	username := req.URL.Query().Get("username")
	if username == "" {
		helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: "username not provided"})
		return
	}

	post := &domain.Post{Conn: ph.Conn}

	posts, err := post.FindByUsername(username)
	if err != nil {
		log.WithError(err).Error("can't find posts")

		helpers.Response(res, http.StatusBadRequest, helpers.Error{Error: err.Error()})
		return
	}

	helpers.Response(res, http.StatusOK, posts)
	return
}
