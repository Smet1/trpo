package config

import "testing"

func TestReadConfig(t *testing.T) {
	type args struct {
		fileName string
		config   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test invalid file name",
			args: args{
				fileName: "kek_filename",
				config:   Config{},
			},
			wantErr: true,
		},
		{
			name: "test valid file",
			args: args{
				fileName: "./../../configs/config.yaml",
				config:   &Config{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadConfig(tt.args.fileName, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
