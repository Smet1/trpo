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

func (p *Post) Insert(db *sqlx.DB) error {
	p.Created = time.Now()
	query := `
INSERT INTO posts (header, short_topic, main_topic, user_id, show, created) 
VALUES (:header, :short_topic, :main_topic, :user_id, :show, :created)
RETURNING id
`
	row, err := db.NamedQuery(query, p)
	if err != nil {
		return errors.Wrap(err, "can't do query")
	}
	defer row.Close()

	res := &sql.NullInt64{}
	for row.Next() {
		err = row.Scan(res)
		if err != nil {
			return errors.Wrap(err, "can't get id")
		}
	}

	p.ID = res.Int64

	return nil
}

type Posts struct {
	Posts []*Post
}

func (p *Posts) Select(db *sqlx.DB, userID int64) error {
	query := `
SELECT (id, header, short_topic, main_topic, user_id, show, created)
FROM posts 
WHERE user_id = $1
`
	err := db.Select(p.Posts, query)
	if err != nil {
		return errors.Wrap(err, "can't do query")
	}

	return nil
}
