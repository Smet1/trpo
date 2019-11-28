package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Tag struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func (t *Tag) Insert(db *sqlx.DB) error {
	query := `
INSERT INTO tag (name) 
VALUES (:name)
RETURNING id
`
	row, err := db.NamedQuery(query, t)
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

	t.ID = res.Int64

	return nil
}

func (t *Tag) LinkWithPost(db *sqlx.DB, postID int64) error {
	query := `
INSERT INTO post_tags (post_id, tag_id) 
VALUES ($1, $2)
`
	_, err := db.Query(query, postID, t.ID)
	if err != nil {
		return errors.Wrap(err, "can't link post with tag")
	}

	return nil
}

func (t *Tag) FindByName(db *sqlx.DB, name string) error {
	query := `
SELECT id, name
FROM tag 
WHERE name = $1
`
	err := db.Get(t, query, name)
	if err != nil {
		return errors.Wrap(err, "can't do query")
	}

	return nil
}

type Tags struct {
	Tags []*Tag
}

func (ts *Tags) GetTagsByPostID(db *sqlx.DB, postID int64) error {
	query := `
SELECT t.id, t.name
FROM post_tags
LEFT JOIN tag t on post_tags.tag_id = t.id
WHERE post_id = $1
`
	err := db.Select(ts.Tags, query)
	if err != nil {
		return errors.Wrap(err, "can't do query")
	}

	return nil
}
