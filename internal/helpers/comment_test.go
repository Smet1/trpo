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
		name    string
		fields  fields
		args    args
		wantErr bool
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
		})
	}
}
