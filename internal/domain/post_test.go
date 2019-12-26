package domain

import (
	"reflect"
	"testing"
	"time"

	"github.com/Smet1/trpo/internal/helpers"
)

func TestPost_ToResponse(t *testing.T) {
	type fields struct {
		ID         int64
		Header     string
		ShortTopic string
		MainTopic  string
		Username   string
		Tags       []string
		Show       bool
		Created    time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *helpers.Post
	}{
		{
			name: "not empty",
			fields: fields{
				ID:         1,
				Header:     "kek_header",
				ShortTopic: "kek_shortTopic",
				MainTopic:  "kek_mainTopic",
				Username:   "kek_username",
				Tags:       []string{"kek", "kek2"},
				Show:       true,
				Created:    time.Unix(1000000000001, 0),
			},
			want: &helpers.Post{
				ID:         1,
				Header:     "kek_header",
				ShortTopic: "kek_shortTopic",
				MainTopic:  "kek_mainTopic",
				Username:   "kek_username",
				Tags:       []string{"kek", "kek2"},
				Show:       true,
				Created:    time.Unix(1000000000001, 0),
			},
		},
		{
			name: "empty",
			fields: fields{
				ID:         0,
				Header:     "",
				ShortTopic: "",
				MainTopic:  "",
				Username:   "",
				Tags:       nil,
				Show:       false,
				Created:    time.Time{},
			},
			want: &helpers.Post{
				ID:         0,
				Header:     "",
				ShortTopic: "",
				MainTopic:  "",
				Username:   "",
				Show:       false,
				Tags:       nil,
				Created:    time.Time{},
			},
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
				Tags:       tt.fields.Tags,
				Show:       tt.fields.Show,
				Created:    tt.fields.Created,
			}
			if got := p.ToResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
