package domain

import (
	"reflect"
	"testing"
	"time"

	"github.com/Smet1/trpo/internal/helpers"
)

func TestComment_ToResponse(t *testing.T) {
	type fields struct {
		ID       int64
		ParentID int64
		Username string
		PostID   int64
		Payload  string
		Show     bool
		Created  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *helpers.Comment
	}{
		{
			name: "empty",
			fields: fields{
				ID:       0,
				ParentID: 0,
				Username: "",
				PostID:   0,
				Payload:  "",
				Show:     false,
				Created:  time.Time{},
			},
			want: &helpers.Comment{
				ID:       0,
				ParentID: 0,
				Username: "",
				PostID:   0,
				Payload:  "",
				Show:     false,
				Created:  time.Time{},
			},
		},
		{
			name: "not empty",
			fields: fields{
				ID:       1,
				ParentID: 1,
				Username: "kek",
				PostID:   1,
				Payload:  "1",
				Show:     true,
				Created:  time.Unix(100000000000, 0),
			},
			want: &helpers.Comment{
				ID:       1,
				ParentID: 1,
				Username: "kek",
				PostID:   1,
				Payload:  "1",
				Show:     true,
				Created:  time.Unix(100000000000, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Comment{
				ID:       tt.fields.ID,
				ParentID: tt.fields.ParentID,
				Username: tt.fields.Username,
				PostID:   tt.fields.PostID,
				Payload:  tt.fields.Payload,
				Show:     tt.fields.Show,
				Created:  tt.fields.Created,
			}
			if got := c.ToResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
