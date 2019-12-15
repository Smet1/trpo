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

type Comments struct {
	Comments []*Comment
}

type CommentTable struct {
	db *sqlx.DB
}

func (ct *CommentTable) Insert(parentID, userID, postID int64, payload string, show bool) (*Comment, error) {
	c := &Comment{
		ParentID: parentID,
		UserID:   userID,
		PostID:   postID,
		Payload:  payload,
		Show:     show,
		Created:  time.Now(),
	}

	query := `
INSERT INTO comments (parent_id, user_id, post_id, payload, show, created) 
VALUES (:parent_id, :user_id, :post_id, :payload, :show, :created)
RETURNING id
`
	row, err := ct.db.NamedQuery(query, c)
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

	c.ID = res.Int64

	return c, nil
}

func (ct *CommentTable) GetCommentsByUserID(userID int64) (*Comments, error) {
	cs := &Comments{}

	query := `
SELECT id, parent_id, user_id, post_id, payload, show, created
FROM comments 
WHERE user_id = $1
`
	err := ct.db.Select(&cs.Comments, query, userID)
	if err != nil {
		return nil, errors.Wrap(err, "can't do query")
	}

	return cs, nil
}

func (ct *CommentTable) GetCommentsByUsername(username string) (*Comments, error) {
	cs := &Comments{}

	query := `
SELECT c.id, c.parent_id, c.user_id, c.post_id, c.payload, c.show, c.created
FROM comments AS c
         JOIN users u on c.user_id = u.id
WHERE u.login = $1 
`
	err := ct.db.Select(&cs.Comments, query, username)
	if err != nil {
		return nil, errors.Wrap(err, "can't do query")
	}

	return cs, nil
}
