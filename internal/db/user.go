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

type UserTable struct {
	db *sqlx.DB
}

func (ut *UserTable) Insert(login, password, avatar string, karma float64) (*User, error) {
	u := &User{
		Login:      login,
		Password:   password,
		Avatar:     avatar,
		Karma:      karma,
		Registered: time.Now(),
	}

	query := `
INSERT INTO users (login, password, avatar, karma, registered) 
VALUES (:login, :password, :avatar, :karma, :registered)
RETURNING id
`
	row, err := ut.db.NamedQuery(query, u)
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

	u.ID = res.Int64

	return u, nil
}

func (ut *UserTable) GetUserByLogin(login string) (*User, error) {
	u := &User{}

	query := `
SELECT id, login, password, avatar, karma, registered
FROM users 
WHERE login = $1
`
	err := ut.db.Get(u, query, login)
	if err != nil {
		return nil, errors.Wrap(err, "can't do query")
	}

	return u, nil
}

func (ut *UserTable) GetUserByID(id int64) (*User, error) {
	u := &User{}

	query := `
SELECT id, login, password, avatar, karma, registered
FROM users 
WHERE id = $1
`
	err := ut.db.Get(ut, query, id)
	if err != nil {
		return nil, errors.Wrap(err, "can't do query")
	}

	return u, nil
}
