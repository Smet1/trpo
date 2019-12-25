package helpers

import (
	"io"
	"strings"
	"testing"
	"time"
)

func TestUser_ParseFromRequest(t *testing.T) {
	type fields struct {
		ID         int64
		Login      string
		Password   string
		Avatar     string
		Karma      float64
		Registered time.Time
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
			name: "test invalid",
			fields: fields{
				ID:         0,
				Login:      "",
				Password:   "",
				Avatar:     "",
				Karma:      0,
				Registered: time.Time{},
			},
			args: args{
				requestBody: strings.NewReader("{kek}"),
			},
			wantErr: true,
		},
		{
			name: "test valid",
			fields: fields{
				ID:         0,
				Login:      "",
				Password:   "",
				Avatar:     "",
				Karma:      0,
				Registered: time.Time{},
			},
			args: args{
				requestBody: strings.NewReader(`{"kek": ""}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:         tt.fields.ID,
				Login:      tt.fields.Login,
				Password:   tt.fields.Password,
				Avatar:     tt.fields.Avatar,
				Karma:      tt.fields.Karma,
				Registered: tt.fields.Registered,
			}
			if err := u.ParseFromRequest(tt.args.requestBody); (err != nil) != tt.wantErr {
				t.Errorf("ParseFromRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
