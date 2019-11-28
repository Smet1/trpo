package domain

import (
	"time"

	"github.com/Smet1/trpo/internal/helpers"

	"github.com/Smet1/trpo/internal/db"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type User struct {
	ID         int64
	Login      string
	Password   string
	Avatar     string
	Karma      float64
	Registered time.Time
	Conn       *sqlx.DB
}

func (u *User) FromParsedRequest(parsed *helpers.User) {
	u.ID = parsed.ID
	u.Login = parsed.Login
	u.Password = parsed.Password
	u.Avatar = parsed.Avatar
	u.Karma = parsed.Karma
	u.Registered = parsed.Registered
}

func (u *User) ToResponse() *helpers.User {
	return &helpers.User{
		ID:         u.ID,
		Login:      u.Login,
		Password:   u.Password,
		Avatar:     u.Avatar,
		Karma:      u.Karma,
		Registered: u.Registered,
	}
}

func (u *User) Validate() error {
	switch {
	case len(u.Login) < 3:
		return errors.New("login len < 3")
	case len(u.Password) < 3:
		return errors.New("password len < 3")
	}

	return nil
}

func (u *User) Create() error {
	userDB := db.User{
		ID:         u.ID,
		Login:      u.Login,
		Password:   u.Password,
		Avatar:     u.Avatar,
		Karma:      u.Karma,
		Registered: u.Registered,
	}

	err := userDB.Insert(u.Conn)
	if err != nil {
		return errors.Wrap(err, "can't create user")
	}

	u.ID = userDB.ID
	u.Registered = userDB.Registered

	return nil
}

func (u *User) GetByUsername(username string) error {
	userDB := db.User{}

	err := userDB.GetUserByLogin(u.Conn, username)
	if err != nil {
		return errors.Wrap(err, "can't find user")
	}

	u.Login = username
	u.Karma = userDB.Karma
	u.Registered = userDB.Registered

	return nil
}

func (u *User) Find(login string) error {
	userDB := db.User{}

	err := userDB.GetUserByLogin(u.Conn, login)
	if err != nil {
		return errors.Wrap(err, "can't find user")
	}

	u.ID = userDB.ID
	u.Login = login
	u.Password = userDB.Password
	u.Karma = userDB.Karma
	u.Registered = userDB.Registered

	return nil
}

func (u *User) Auth(login, password string) (int64, error) {
	err := u.Find(login)
	if err != nil {
		return -1, errors.Wrap(err, "user not found")
	}

	if u.Password != password {
		return -1, errors.Wrap(err, "wrong password")
	}

	return u.ID, nil
}
