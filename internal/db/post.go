package db

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Post struct {
	ID         int64     `db:"id"`
	Header     string    `db:"header"`
	ShortTopic string    `db:"short_topic"`
	MainTopic  string    `db:"main_topic"`
	UserID     int64     `db:"user_id"`
	Show       bool      `db:"show"`
	Created    time.Time `db:"created"`
}

type Posts struct {
	Posts []*Post
}

type PostTable struct {
	db *sqlx.DB
}

func (pt *PostTable) Insert(header, shortTopic, mainTopic string, userID int64, show bool) (*Post, error) {
	p := &Post{
		Header:     header,
		ShortTopic: shortTopic,
		MainTopic:  mainTopic,
		UserID:     userID,
		Show:       show,
		Created:    time.Now(),
	}

	query := `
INSERT INTO posts (header, short_topic, main_topic, user_id, show, created) 
VALUES (:header, :short_topic, :main_topic, :user_id, :show, :created)
RETURNING id
`
	row, err := pt.db.NamedQuery(query, p)
	if err != nil {
		return nil, errors.Wrap(err, "can't do query")
	}
	defer row.Close()

	res := &sql.NullInt64{}
	for row.Next() {
		err = row.Scan(res)
		if err != nil {
			return nil, errors.Wrap(err, "can't get id")
		}
	}

	p.ID = res.Int64

	return p, nil
}

func (pt *PostTable) GetPostByID(id int64) ([]*Post, error) {
	p := &Posts{}

	query := `
SELECT id, header, short_topic, main_topic, user_id, show, created
FROM posts 
WHERE id = $1
`
	err := pt.db.Get(p, query, id)
	if err != nil {
		return nil, errors.Wrap(err, "can't do query")
	}

	return p.Posts, nil
}

func (pt *PostTable) GetPostsByUserID(userID int64) ([]*Post, error) {
	p := &Posts{}

	query := `
SELECT id, header, short_topic, main_topic, user_id, show, created
FROM posts 
WHERE user_id = $1
`
	err := pt.db.Select(&p.Posts, query, userID)
	if err != nil {
		return nil, errors.Wrap(err, "can't do query")
	}

	return p.Posts, nil
}
