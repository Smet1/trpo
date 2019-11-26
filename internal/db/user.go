package db

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type User struct {
	ID         int64     `db:"id"`
	Login      string    `db:"login"`
	Password   string    `db:"password"`
	Avatar     string    `db:"avatar"`
	Karma      float64   `db:"karma"`
	Registered time.Time `db:"registered"`
}

func (u *User) Insert(db *sqlx.DB) error {
	u.Registered = time.Now()
	query := `
INSERT INTO users (login, password, avatar, karma, registered) 
VALUES (:login, :password, :avatar, :karma, :registered)
RETURNING id
`
	row, err := db.NamedQuery(query, u)
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

	u.ID = res.Int64

	return nil
}

func (u *User) Select(db *sqlx.DB, login string) error {
	u.Registered = time.Now()
	query := `
SELECT id, login, password, avatar, karma, registered
FROM users 
WHERE login = $1
`
	err := db.Get(u, query, login)
	if err != nil {
		return errors.Wrap(err, "can't do query")
	}

	return nil
}
