package helpers

import (
	"io"
	"strings"
	"testing"
	"time"
)

func TestPost_ParseFromRequest(t *testing.T) {
	type fields struct {
		ID         int64
		Header     string
		ShortTopic string
		MainTopic  string
		Username   string
		Show       bool
		Tags       []string
		Created    time.Time
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
				ID:         0,
				Header:     "",
				ShortTopic: "",
				MainTopic:  "",
				Username:   "",
				Show:       false,
				Tags:       nil,
				Created:    time.Time{},
			},
			args: args{
				requestBody: strings.NewReader("{kek}"),
			},
			wantErr: true,
		},
		{
			name: "test invalid input",
			fields: fields{
				ID:         0,
				Header:     "",
				ShortTopic: "",
				MainTopic:  "",
				Username:   "",
				Show:       false,
				Tags:       nil,
				Created:    time.Time{},
			},
			args: args{
				requestBody: strings.NewReader(`{"kek": ""}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				ID:         tt.fields.ID,
				Header:     tt.fields.Header,
				ShortTopic: tt.fields.ShortTopic,
				MainTopic:  tt.fields.MainTopic,
				Username:   tt.fields.Username,
				Show:       tt.fields.Show,
				Tags:       tt.fields.Tags,
				Created:    tt.fields.Created,
			}
			if err := p.ParseFromRequest(tt.args.requestBody); (err != nil) != tt.wantErr {
				t.Errorf("ParseFromRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
