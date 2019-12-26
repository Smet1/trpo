package helpers

import (
	"io"
	"strings"
	"testing"
	"time"
)

func TestComment_ParseFromRequest(t *testing.T) {
	type fields struct {
		ID       int64
		ParentID int64
		Username string
		PostID   int64
		Payload  string
		Show     bool
		Created  time.Time
	}
	type args struct {
		requestBody io.Reader
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantFields fields
	}{
		{
			name: "test invalid input",
			fields: fields{
				ID:       0,
				ParentID: 0,
				Username: "",
				PostID:   0,
				Payload:  "",
				Show:     false,
				Created:  time.Time{},
			},
			args: args{
				requestBody: strings.NewReader(`{kek}`),
			},
			wantErr: true,
		},
		{
			name: "test invalid input",
			fields: fields{
				ID:       0,
				ParentID: 0,
				Username: "",
				PostID:   0,
				Payload:  "",
				Show:     false,
				Created:  time.Time{},
			},
			args: args{
				requestBody: strings.NewReader(`{"kek": "",}`),
			},
			wantErr: true,
		},
		{
			name: "test ok input",
			fields: fields{
				ID:       0,
				ParentID: 0,
				Username: "",
				PostID:   0,
				Payload:  "",
				Show:     false,
				Created:  time.Time{},
			},
			args: args{
				requestBody: strings.NewReader(`{"kek": ""}`),
			},
			wantErr: false,
		},
		{
			name: "test ok input",
			fields: fields{
				ID:       0,
				ParentID: 0,
				Username: "",
				PostID:   0,
				Payload:  "",
				Show:     false,
				Created:  time.Time{},
			},
			args: args{
				requestBody: strings.NewReader(`{"id": 100, "parent_id": 101, "username": "kek", "post_id": 102, "payload": "kek", "show": true}`),
			},
			wantErr: false,
			wantFields: fields{
				ID:       100,
				ParentID: 101,
				Username: "kek",
				PostID:   102,
				Payload:  "kek",
				Show:     true,
				Created:  time.Time{},
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
			if err := c.ParseFromRequest(tt.args.requestBody); (err != nil) != tt.wantErr {
				t.Errorf("ParseFromRequest() error = %v, wantErr %v", err, tt.wantErr)
			}

			switch {
			case c.Show != tt.wantFields.Show:
				fallthrough
			case c.ID != tt.wantFields.ID:
				fallthrough
			case c.Created != tt.wantFields.Created:
				fallthrough
			case c.ParentID != tt.wantFields.ParentID:
				fallthrough
			case c.Payload != tt.wantFields.Payload:
				fallthrough
			case c.PostID != tt.wantFields.PostID:
				fallthrough
			case c.Username != tt.wantFields.Username:
				t.Errorf("ToResponse() = %v, want %v", c, tt.wantFields)
			}
		})
	}
}
