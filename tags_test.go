package main

import (
	"reflect"
	"testing"
)

func Test_tags_ReadTagsFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		t       *tags
		args    args
		want    []string
		wantErr bool
	}{
		{
			"should return 0 tags as string array",
			&tags{},
			struct{ file string }{"testdata/tagsData0"},
			[]string{},
			false,
		},
		{
			"should return 2 tags as string array",
			&tags{},
			struct{ file string }{"testdata/tagsData1"},
			[]string{"0.0.0", "latest"},
			false,
		},
		{
			"should return 1 tags as string array",
			&tags{},
			struct{ file string }{"testdata/tagsData2"},
			[]string{"0.0.0"},
			false,
		},
		{
			"should return error",
			&tags{},
			struct{ file string }{"testdata/tagsData"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.ReadTagsFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("tags.ReadTagsFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tags.ReadTagsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tags_GetTags(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		t    *tags
		args args
		want []string
	}{
		{
			"should return 1 tag with jobnumber 1",
			&tags{"testdata/tagsData0"},
			struct{ config Config }{Config{
				UseGitTag: false,
				JobNum:    "1",
			}},
			[]string{"latest-1"},
		},
		{
			"should return 2 tags with jobnumber 1",
			&tags{"testdata/tagsData1"},
			struct{ config Config }{Config{
				UseGitTag: false,
				JobNum:    "1",
			}},
			[]string{"0.0.0-1", "latest-1"},
		},
		{
			"should return 2 tags with jobnumber 2",
			&tags{"testdata/tagsData1"},
			struct{ config Config }{Config{
				UseGitTag: false,
				JobNum:    "2",
			}},
			[]string{"0.0.0-2", "latest-2"},
		},
		{
			"should return 1 tags with jobnumber 1",
			&tags{"testdata/tagsData2"},
			struct{ config Config }{Config{
				UseGitTag: false,
				JobNum:    "1",
			}},
			[]string{"0.0.0-1"},
		},
		{
			"should return 'gitTag' and 'latest' without jobnumber",
			&tags{"testdata/tagsData1"},
			struct{ config Config }{Config{
				UseGitTag: true,
				GitTag:    "1.1.1",
				JobNum:    "1",
			}},
			[]string{"1.1.1", "latest"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.GetTags(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tags.GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTags(t *testing.T) {
	tests := []struct {
		name string
		want Tags
	}{
		{
			"should create new tags struct with .tags file",
			&tags{".tags"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
