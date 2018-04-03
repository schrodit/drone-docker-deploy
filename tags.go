package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Tags interface {
	GetTags(Config) []string
	ReadTagsFile(string) ([]string, error)
}

type tags struct {
	tagsFile string
}

func NewTags() Tags {
	return &tags{".tags"}
}

func (t *tags) ReadTagsFile(file string) ([]string, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if len(dat) == 0 {
		return []string{}, nil
	}

	versionsString := string(dat)
	versionsString = strings.Replace(versionsString, " ", "", -1)
	versions := strings.Split(versionsString, ",")
	return versions, nil
}

func (t *tags) GetTags(config Config) []string {
	if config.UseGitTag == true {
		return []string{config.GitTag, "latest"}
	}

	tags, err := t.ReadTagsFile(t.tagsFile)
	if err != nil || len(tags) == 0 {
		tags = []string{"latest"}
	}
	for i, tag := range tags {
		if tag != "latest" {
			tags[i] = fmt.Sprintf("%s-%s", tag, config.JobNum)
		}
	}
	return tags
}
