package domain

import (
	"time"

	"github.com/Smet1/trpo/internal/helpers"
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

func (c *Comment) FromParsedRequest(parsed *helpers.Comment) {
	c.ID = parsed.ID
	c.ParentID = parsed.ParentID
	c.Username = parsed.Username
	c.PostID = parsed.PostID
	c.Payload = parsed.Payload
	c.Show = parsed.Show
	c.Created = parsed.Created
}

func (c *Comment) ToResponse() *helpers.Comment {
	return &helpers.Comment{
		ID:       c.ID,
		ParentID: c.ParentID,
		Username: c.Username,
		PostID:   c.PostID,
		Payload:  c.Payload,
		Show:     c.Show,
		Created:  c.Created,
	}
}
