package domain

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"

	"github.com/Smet1/trpo/internal/db"
	"github.com/Smet1/trpo/internal/helpers"
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

func (p *Post) Create(conn *sqlx.DB) error {
	userTable := db.UserTable{Conn: conn}
	user, err := userTable.GetUserByLogin(p.Username)
	if err != nil {
		return errors.Wrap(err, "can't find post's creator")
	}

	postTable := db.PostTable{Conn: conn}
	post, err := postTable.Insert(p.Header, p.ShortTopic, p.MainTopic, user.ID, p.Show)
	if err != nil {
		return errors.Wrap(err, "can't insert post")
	}

	p.ID = post.ID
	p.Created = post.Created

	// tagDB := &db.Tag{}
	tagsTable := &db.TagsTable{Conn: conn}
	postTagsTable := &db.PostTagsTable{Conn: conn}

	for _, tag := range p.Tags {
		foundTag, err := tagsTable.FindByName(tag)
		if err != nil {
			return errors.Wrap(err, "can't find tag")
		}

		err = postTagsTable.LinkWithPost(p.ID, foundTag.ID)
		if err != nil {
			return errors.Wrap(err, "can't link tag with post")
		}
	}

	return nil
}

func (p *Post) FindByID(postID int64, conn *sqlx.DB) (*helpers.Post, error) {
	postTable := db.PostTable{Conn: conn}
	post, err := postTable.GetPostByID(postID)
	if err != nil {
		return nil, errors.Wrap(err, "can't find post")
	}

	tagsTable := db.PostTagsTable{Conn: conn}
	tags, err := tagsTable.GetTagsByPostID(postID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrap(err, "can't find post tags")
		}
	}
	tagNames := make([]string, 0, len(tags))

	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}

	userTable := db.UserTable{Conn: conn}
	user, err := userTable.GetUserByID(post.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "can't find post's creator")
	}

	return &helpers.Post{
		ID:         post.ID,
		Header:     post.Header,
		ShortTopic: post.ShortTopic,
		MainTopic:  post.MainTopic,
		Username:   user.Login,
		Tags:       tagNames,
		Show:       post.Show,
		Created:    post.Created,
	}, nil
}

func (p *Post) FindByUsername(username string, conn *sqlx.DB) ([]*helpers.Post, error) {
	userTable := db.UserTable{Conn: conn}
	user, err := userTable.GetUserByLogin(username)
	if err != nil {
		return nil, errors.Wrap(err, "can't find post's creator")
	}

	postTable := db.PostTable{Conn: conn}
	postsDB, err := postTable.GetPostsByUserID(user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "can't find posts")
	}

	posts := make([]*helpers.Post, 0, len(postsDB))

	postTagsTable := db.PostTagsTable{Conn: conn}
	for _, post := range postsDB {
		tags, err := postTagsTable.GetTagsByPostID(post.ID)
		if err != nil {
			return nil, errors.Wrap(err, "can't find post tags")
		}

		tagsResp := make([]string, 0, len(tags))
		for _, tag := range tags {
			tagsResp = append(tagsResp, tag.Name)
		}
		posts = append(posts, &helpers.Post{
			ID:         post.ID,
			Header:     post.Header,
			ShortTopic: post.ShortTopic,
			MainTopic:  post.MainTopic,
			Username:   user.Login,
			Show:       post.Show,
			Tags:       tagsResp,
			Created:    post.Created,
		})
	}

	return posts, nil
}
