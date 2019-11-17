package db

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

type User struct {
	ID         int       `db:"id"`
	Login      string    `db:"login"`
	Password   string    `db:"password"`
	Avatar     string    `db:"avatar"`
	Karma      float64   `db:"karma"`
	Registered time.Time `db:"registered"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(&u,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&u.Login, validation.Required, validation.Length(5, 50)),
	)
}

func (u *User) Insert() {}

func (u *User) Select() {

}

func (u *User) Update() {

}

func (u *User) Delete() {

}
