package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
)

type Comment struct {
	ID       int64     `json:"id,omitempty"`
	ParentID int64     `json:"parent_id,omitempty"`
	Username string    `json:"username,omitempty"`
	PostID   int64     `json:"post_id,omitempty"`
	Payload  string    `json:"payload,omitempty"`
	Show     bool      `json:"show,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

func (c *Comment) ParseFromRequest(requestBody io.Reader) error {
	body, _ := ioutil.ReadAll(requestBody)
	err := json.Unmarshal(body, c)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal request body")
	}

	return nil
}
