package domain

import (
	"reflect"
	"testing"
	"time"

	"github.com/Smet1/trpo/internal/helpers"
)

func TestUser_Validate(t *testing.T) {
	type fields struct {
		ID         int64
		Login      string
		Password   string
		Avatar     string
		Karma      float64
		Registered time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok test",
			fields: fields{
				Login:    "aaaaa",
				Password: "wwwww",
			},
			wantErr: false,
		},
		{
			name: "fail test",
			fields: fields{
				Login:    "aa",
				Password: "ww",
			},
			wantErr: true,
		},
		{
			name: "fail test",
			fields: fields{
				Login:    "a\n",
				Password: "w\n",
			},
			wantErr: true,
		},
		{
			name: "fail test",
			fields: fields{
				Login:    "①②③",
				Password: "④⑤⑥",
			},
			wantErr: false,
		},
		{
			name: "short login",
			fields: fields{
				Login:    "①②",
				Password: "④⑤⑥",
			},
			wantErr: true,
		},
		{
			name: "short password",
			fields: fields{
				Login:    "①②③",
				Password: "④⑤",
			},
			wantErr: true,
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
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_ToResponse(t *testing.T) {
	type fields struct {
		ID         int64
		Login      string
		Password   string
		Avatar     string
		Karma      float64
		Registered time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *helpers.User
	}{
		{
			name: "empty",
			fields: fields{
				ID:         0,
				Login:      "",
				Password:   "",
				Avatar:     "",
				Karma:      0,
				Registered: time.Time{},
			},
			want: &helpers.User{},
		},
		{
			name: "not empty",
			fields: fields{
				ID:         1,
				Login:      "kek_login",
				Password:   "kek_password",
				Avatar:     "kek_avatar",
				Karma:      100,
				Registered: time.Unix(100000000000, 0),
			},
			want: &helpers.User{
				ID:         1,
				Login:      "kek_login",
				Password:   "kek_password",
				Avatar:     "kek_avatar",
				Karma:      100,
				Registered: time.Unix(100000000000, 0),
			},
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
			if got := u.ToResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
