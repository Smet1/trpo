package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
)

type Post struct {
	ID         int64     `json:"id,omitempty"`
	Header     string    `json:"header,omitempty"`
	ShortTopic string    `json:"short_topic,omitempty"`
	MainTopic  string    `json:"main_topic,omitempty"`
	Username   string    `json:"username,omitempty"`
	Show       bool      `json:"show,omitempty"`
	Created    time.Time `json:"created,omitempty"`
}

func (p *Post) ParseFromRequest(requestBody io.Reader) error {
	body, _ := ioutil.ReadAll(requestBody)
	err := json.Unmarshal(body, p)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal request body")
	}

	return nil
}
