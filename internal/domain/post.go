package domain

import (
	"database/sql"
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
	Tags       []string
	Show       bool
	Created    time.Time

	Conn *sqlx.DB
}

func (p *Post) FromParsedRequest(parsed *helpers.Post) {
	p.ID = parsed.ID
	p.Header = parsed.Header
	p.ShortTopic = parsed.ShortTopic
	p.MainTopic = parsed.MainTopic
	p.Username = parsed.Username
	p.Tags = parsed.Tags
	p.Show = parsed.Show
}

func (p *Post) ToResponse() *helpers.Post {
	return &helpers.Post{
		ID:         p.ID,
		Header:     p.Header,
		ShortTopic: p.ShortTopic,
		MainTopic:  p.MainTopic,
		Username:   p.Username,
		Tags:       p.Tags,
		Show:       p.Show,
		Created:    p.Created,
	}
}

func (p *Post) Create() error {
	userDB := db.User{}
	err := userDB.GetUserByLogin(p.Conn, p.Username)
	if err != nil {
		return errors.Wrap(err, "can't find post's creator")
	}

	postDB := &db.Post{
		Header:     p.Header,
		ShortTopic: p.ShortTopic,
		MainTopic:  p.MainTopic,
		UserID:     userDB.ID,
		Show:       p.Show,
	}

	err = postDB.Insert(p.Conn)
	if err != nil {
		return errors.Wrap(err, "can't insert post")
	}

	p.ID = postDB.ID
	p.Created = postDB.Created

	tagDB := &db.Tag{}
	validTags := make([]string, 0, len(p.Tags))
	tagIDs := make([]int64, 0, len(p.Tags))
	for _, tag := range p.Tags {
		err := tagDB.FindByName(p.Conn, tag)
		if err != nil {
			return errors.Wrap(err, "can't find tag")
		}

		err = tagDB.LinkWithPost(p.Conn, p.ID)
		if err != nil {
			return errors.Wrap(err, "can't link tag with post")
		}

		validTags = append(validTags, tagDB.Name)
		tagIDs = append(tagIDs, tagDB.ID)
	}

	return nil
}

func (p *Post) FindByID(id int64) (*helpers.Post, error) {
	postDB := db.Post{}
	err := postDB.GetPostByID(p.Conn, id)
	if err != nil {
		return nil, errors.Wrap(err, "can't find post")
	}

	tagsDB := db.Tags{}
	err = tagsDB.GetTagsByPostID(p.Conn, id)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "can't find post tags")
		}
	}
	tags := make([]string, 0, len(tagsDB.Tags))

	for _, tag := range tagsDB.Tags {
		tags = append(tags, tag.Name)
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
		Tags:       tags,
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
	tagsDB := db.Tags{}

	for _, post := range postDB.Posts {
		err = tagsDB.GetTagsByPostID(p.Conn, post.ID)
		if err != nil {
			return nil, errors.Wrap(err, "can't find post tags")
		}

		tags := make([]string, 0, len(tagsDB.Tags))
		for _, tag := range tagsDB.Tags {
			tags = append(tags, tag.Name)
		}
		posts = append(posts, &helpers.Post{
			ID:         post.ID,
			Header:     post.Header,
			ShortTopic: post.ShortTopic,
			MainTopic:  post.MainTopic,
			Username:   userDB.Login,
			Show:       post.Show,
			Tags:       tags,
			Created:    post.Created,
		})
	}

	return posts, nil
}
