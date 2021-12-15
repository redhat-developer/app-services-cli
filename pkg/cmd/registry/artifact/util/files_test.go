package util

import (
	"context"
	"reflect"
	"testing"
)

func TestIsURL(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should return true for url starting with https",
			args: args{
				path: "https://bu98.serviceregistry-stage.rhcloud.com/t/8ecff228-1ffe-4cf5-b38b-55223885ee00/apis/registry/v2",
			},
			want: true,
		},
		{
			name: "Should return true for url starting with http",
			args: args{
				path: "http://localhost:8082/",
			},
			want: true,
		},
		{
			name: "Should return false for regular file path",
			args: args{
				path: "./schema/artifact.json",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsURL(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsURL(%v) = %v, want %v", tt.args.path, got, tt.want)
			}
		})
	}
}

func TestGetContentFromFileURL(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should load file content from valid url",
			args: args{
				path: "https://raw.githubusercontent.com/bolcom/avro-schema-viewer/master/docs/assets/avsc/1.0/schema.avsc",
			},
			wantErr: false,
		},
		{
			name: "Should throw error if URL is not found",
			args: args{
				path: "https://test-123-test-404.com/test",
			},
			wantErr: true,
		},
		{
			name: "Should throw error if argument is not a URL",
			args: args{
				path: "./schema/artifact.json",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetContentFromFileURL(context.TODO(), tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("GetContentFromFileURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
