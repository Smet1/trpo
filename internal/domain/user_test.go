package domain

import (
	"testing"
	"time"
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
