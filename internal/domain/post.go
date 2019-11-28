package domain

import (
	"time"

	"github.com/Smet1/trpo/internal/db"
	"github.com/Smet1/trpo/internal/helpers"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Post struct {
	ID         int64
	Header     string
	ShortTopic string
	MainTopic  string
	Username   string
	Show       bool
	Created    time.Time

	User *User
	Conn *sqlx.DB
}

func (p *Post) FromParsedRequest(parsed *helpers.Post) {
	p.ID = parsed.ID
	p.Header = parsed.Header
	p.ShortTopic = parsed.ShortTopic
	p.MainTopic = parsed.MainTopic
	p.Username = parsed.Username
	p.Show = parsed.Show
}

func (p *Post) ToResponse() *helpers.Post {
	return &helpers.Post{
		ID:         p.ID,
		Header:     p.Header,
		ShortTopic: p.ShortTopic,
		MainTopic:  p.MainTopic,
		Username:   p.User.Login,
		Show:       p.Show,
		Created:    p.Created,
	}
}

func (p *Post) FindByID(id int64) (*helpers.Post, error) {
	postDB := db.Post{}
	err := postDB.GetPostByID(p.Conn, id)
	if err != nil {
		return nil, errors.Wrap(err, "can't find post")
	}

	userDB := db.User{}
	err = userDB.GetUserByID(p.Conn, postDB.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "can't find post's creator")
	}

	return &helpers.Post{
		ID:         postDB.ID,
		Header:     postDB.Header,
		ShortTopic: postDB.ShortTopic,
		MainTopic:  postDB.MainTopic,
		Username:   userDB.Login,
		Show:       postDB.Show,
		Created:    postDB.Created,
	}, nil
}

func (p *Post) FindByUsername(username string) ([]*helpers.Post, error) {
	userDB := db.User{}
	err := userDB.GetUserByLogin(p.Conn, username)
	if err != nil {
		return nil, errors.Wrap(err, "can't find post's creator")
	}

	postDB := db.Posts{}
	err = postDB.GetPostsByUserID(p.Conn, userDB.ID)
	if err != nil {
		return nil, errors.Wrap(err, "can't find posts")
	}

	posts := make([]*helpers.Post, 0, len(postDB.Posts))

	for _, post := range postDB.Posts {
		posts = append(posts, &helpers.Post{
			ID:         post.ID,
			Header:     post.Header,
			ShortTopic: post.ShortTopic,
			MainTopic:  post.MainTopic,
			Username:   userDB.Login,
			Show:       post.Show,
			Created:    post.Created,
		})
	}

	return posts, nil
}
