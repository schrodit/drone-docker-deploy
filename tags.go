package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadTagsFile(file string) ([]string, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	versionsString := string(dat)
	versionsString = strings.Replace(versionsString, " ", "", -1)
	versions := strings.Split(versionsString, ",")
	return versions, nil
}

func GetTags(config Config) []string {
	if config.UseGitTag == true {
		return []string{config.GitTag}
	}

	tags, err := ReadTagsFile(".tags")
	if err != nil || len(tags) == 0 {
		tags = []string{"latest"}
	}
	for i, tag := range tags {
		fmt.Println(tag)
		tags[i] = fmt.Sprintf("%s-%s", tag, config.JobNum)
	}
	return tags
}
