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

func (u *User) Create(conn *sqlx.DB) error {
	userTable := db.UserTable{Conn: conn}

	user, err := userTable.Insert(u.Login, u.Password, u.Avatar, u.Karma)
	if err != nil {
		return errors.Wrap(err, "can't create user")
	}

	u.ID = user.ID
	u.Registered = user.Registered

	return nil
}

func (u *User) GetByUsername(username string, conn *sqlx.DB) error {
	userTable := db.UserTable{Conn: conn}

	user, err := userTable.GetUserByLogin(username)
	if err != nil {
		return errors.Wrap(err, "can't find user")
	}

	u.Login = username
	u.Karma = user.Karma
	u.Registered = user.Registered

	return nil
}

func (u *User) Find(login string, conn *sqlx.DB) error {
	userTable := db.UserTable{Conn: conn}

	user, err := userTable.GetUserByLogin(login)
	if err != nil {
		return errors.Wrap(err, "can't find user")
	}

	u.ID = user.ID
	u.Login = login
	u.Password = user.Password
	u.Karma = user.Karma
	u.Registered = user.Registered

	return nil
}

func (u *User) Auth(login, password string, conn *sqlx.DB) (int64, error) {
	err := u.Find(login, conn)
	if err != nil {
		return -1, errors.Wrap(err, "user not found")
	}

	if u.Password != password {
		return -1, errors.Wrap(err, "wrong password")
	}

	return u.ID, nil
}
