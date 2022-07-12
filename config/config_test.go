package config

import (
	"errors"
	"github.com/agiledragon/gomonkey"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		args       args
		wantConfig *Template
		wantErr    bool
	}{
		{
			name: "init global config from file",
			args: args{
				path: "/etc/test/config.yaml",
			},
			wantConfig: &Template{
				WorkDirectory: "/etc/test",
				Web: WebTemplate{
					Host: "127.0.0.1",
					Port: 80,
					Mode: "release",
				},
				Database: DatabaseTemplate{
					MySQL: MySQLTemplate{
						Address:  "127.0.0.1:3306",
						DBName:   "test",
						Username: "test",
						Password: "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "init global config from not exist file",
			args: args{
				path: "",
			},
			wantConfig: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patches := gomonkey.ApplyFunc(load, func(path string) (config *Template, err error) {
				if len(path) == 0 {
					return nil, &os.PathError{Op: "open", Path: path, Err: errors.New("file not exist")}
				}

				return tt.wantConfig, nil
			})
			defer patches.Reset()

			err := Init(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(globalConfig, tt.wantConfig) {
				t.Errorf("Init() globalConfig = %v, want %v", globalConfig, tt.wantConfig)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		want *Template
	}{
		{
			name: "get global config",
			want: &Template{
				WorkDirectory: "/etc/test",
				Web: WebTemplate{
					Host: "127.0.0.1",
					Port: 80,
					Mode: "release",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patches := gomonkey.ApplyGlobalVar(&globalConfig, tt.want)
			defer patches.Reset()

			if got := Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_load(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		args       args
		wantConfig *Template
		wantErr    bool
	}{
		{
			name: "load config from file",
			args: args{
				path: "/etc/test/config.yaml",
			},
			wantConfig: &Template{
				WorkDirectory: "/etc/test",
				Web: WebTemplate{
					Host: "127.0.0.1",
					Port: 80,
					Mode: "release",
				},
				Database: DatabaseTemplate{
					MySQL: MySQLTemplate{
						Address:  "127.0.0.1:3306",
						DBName:   "test",
						Username: "test",
						Password: "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "load config from not exist file",
			args: args{
				path: "",
			},
			wantConfig: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patches := gomonkey.ApplyFunc(ioutil.ReadFile, func(filename string) ([]byte, error) {
				if len(filename) == 0 {
					return nil, &os.PathError{Op: "open", Path: filename, Err: errors.New("file not exist")}
				}

				return yaml.Marshal(tt.wantConfig)
			})
			defer patches.Reset()

			gotConfig, err := load(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("load() gotConfig = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}
