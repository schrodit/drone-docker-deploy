package main

import (
	"io/ioutil"
	"strings"
)

func ReadTagsFile(file string) ([]string, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	versionsString := string(dat)
	versions := strings.Split(versionsString, ",")
	return versions, nil
}
