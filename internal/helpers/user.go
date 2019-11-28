package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
)

type User struct {
	ID         int64     `json:"id,omitempty"`
	Login      string    `json:"login,omitempty"`
	Password   string    `json:"password,omitempty"`
	Avatar     string    `json:"avatar,omitempty"`
	Karma      float64   `json:"karma,omitempty"`
	Registered time.Time `json:"registered,omitempty"`
}

func (u *User) ParseFromRequest(requestBody io.Reader) error {
	body, _ := ioutil.ReadAll(requestBody)
	err := json.Unmarshal(body, u)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal request body")
	}

	return nil
}
