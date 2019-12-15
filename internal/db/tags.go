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

type TagsTable struct {
	Conn *sqlx.DB
}

func (tt *TagsTable) Insert(name string) (*Tag, error) {
	t := &Tag{
		Name: name,
	}

	query := `
INSERT INTO tag (name) 
VALUES (:name)
RETURNING id
`
	row, err := tt.Conn.NamedQuery(query, t)
	if err != nil {
		return t, errors.Wrap(err, "can'tt do query")
	}
	defer row.Close()

	res := &sql.NullInt64{}
	for row.Next() {
		err = row.Scan(res)
		if err != nil {
			return t, errors.Wrap(err, "can'tt get id")
		}
	}

	t.ID = res.Int64

	return t, nil
}

func (tt *TagsTable) FindByName(name string) (*Tag, error) {
	t := &Tag{}

	query := `
SELECT id, name
FROM tag 
WHERE name = $1
`
	err := tt.Conn.Get(t, query, name)
	if err != nil {
		return t, errors.Wrap(err, "can'tt do query")
	}

	return t, nil
}

type Tags struct {
	Tags []*Tag
}

type PostTagsTable struct {
	Conn *sqlx.DB
}

func (ptt *PostTagsTable) LinkWithPost(postID, tagID int64) error {
	query := `
INSERT INTO post_tags (post_id, tag_id) 
VALUES ($1, $2)
`
	_, err := ptt.Conn.Query(query, postID, tagID)
	if err != nil {
		return errors.Wrap(err, "can'tt link post with tag")
	}

	return nil
}

func (ptt *PostTagsTable) GetTagsByPostID(postID int64) ([]*Tag, error) {
	tags := &Tags{}
	query := `
SELECT t.id, t.name
FROM post_tags
LEFT JOIN tag t on post_tags.tag_id = t.id
WHERE post_id = $1
`
	err := ptt.Conn.Select(&tags.Tags, query, postID)
	if err != nil {
		return nil, errors.Wrap(err, "can't do query")
	}

	return tags.Tags, nil
}
