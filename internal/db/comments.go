package db

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Comment struct {
	ID       int64     `db:"id"`
	ParentID int64     `db:"parent_id"`
	UserID   int64     `db:"user_id"`
	PostID   int64     `db:"post_id"`
	Payload  string    `db:"payload"`
	Show     bool      `db:"show"`
	Created  time.Time `db:"created"`
}

func (c *Comment) Insert(db *sqlx.DB) error {
	c.Created = time.Now()
	query := `
INSERT INTO comments (parent_id, user_id, post_id, payload, show, created) 
VALUES (:parent_id, :user_id, :post_id, :payload, :show, :created)
RETURNING id
`
	row, err := db.NamedQuery(query, c)
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

	c.ID = res.Int64

	return nil
}

type Comments struct {
	Comments []*Comment
}

func (cs *Comments) GetByUserID(db *sqlx.DB, userID int64) error {
	query := `
SELECT id, parent_id, user_id, post_id, payload, show, created
FROM comments 
WHERE user_id = $1
`
	err := db.Select(&cs.Comments, query, userID)
	if err != nil {
		return errors.Wrap(err, "can't do query")
	}

	return nil
}

func (cs *Comments) GetByUsername(db *sqlx.DB, username string) error {
	query := `
SELECT c.id, c.parent_id, c.user_id, c.post_id, c.payload, c.show, c.created
FROM comments AS c
         JOIN users u on c.user_id = u.id
WHERE u.login = $1 
`
	err := db.Select(&cs.Comments, query, username)
	if err != nil {
		return errors.Wrap(err, "can't do query")
	}

	return nil
}
