package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Tags interface {
	GetTags(Config) []string
	ReadTagsFile(string) ([]string, error)
	GetNewestGitTag() string
	AddJobNumber([]string, Config) []string
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
	var tags []string
	if config.UseGitTag == true {
		tags = []string{config.GitTag}
	} else {
		var err error
		tags, err = t.ReadTagsFile(t.tagsFile)
		if err != nil || len(tags) == 0 {
			tags = []string{"latest"}
		}
	}

	if config.BuildEvent == "tag" || config.Latest {
		tags = append(tags, "latest")
	}

	if config.AddJobNumber && config.BuildEvent != "tag" {
		tags = t.AddJobNumber(tags, config)
	}

	return tags
}

func (t *tags) AddJobNumber(tags []string, config Config) []string {
	for i, tag := range tags {
		if tag != "latest" {
			tags[i] = fmt.Sprintf("%s-%s", tag, config.JobNum)
		}
	}
	return tags
}

func (t *tags) GetNewestGitTag() string {
	setupGit()

	cmd := exec.Command("git", "fetch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("cannot fetch tags\n%v", err)
		return ""
	}

	cmd = exec.Command("git", "describe", "--tags", "--abbrev=0")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Printf("cannot get latest tag\n%v", err)
		return ""
	}
	tag := strings.Replace(out.String(), "\n", "", -1)

	return tag
}

func setupGit() {
	machine := os.Getenv("DRONE_NETRC_MACHINE")
	username := os.Getenv("DRONE_NETRC_USERNAME")
	password := os.Getenv("DRONE_NETRC_PASSWORD")
	if err := writeNetrc(machine, username, password); err != nil {
		log.Printf("Error writing .netrc file\n%v", err)
	}
}
