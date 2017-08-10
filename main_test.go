package main

import (
	"path/filepath"
	"testing"
)

func Test_getFilename(t *testing.T) {
	type args struct {
		fpath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"with path",
			args{filepath.Join("a", "b", "filename.txt")},
			"filename",
		},
		{
			"just filename with extension",
			args{"filename.txt"},
			"filename",
		},
		{
			"just filename without extension",
			args{"filename"},
			"filename",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFilename(tt.args.fpath); got != tt.want {
				t.Errorf("getFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
