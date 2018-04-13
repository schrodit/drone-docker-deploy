package main

import (
	"fmt"
)

type Config struct {
	Registry     string
	Image        string
	Dockerfile   string
	Dir          string
	Username     string
	Password     string
	Tags         []string
	AddJobNumber bool
	JobNum       string
	Latest       bool
	UseGitTag    bool
	BuildEvent   string
	GitTag       string
}

func main() {
	fmt.Println("starting deployment")
	tags := NewTags()
	envVars := NewEnvVars(tags)
	run(envVars)
	fmt.Println("succesfully published images")
}

func run(e EnvVars) {
	//create Config out of Environment variables
	config := e.Get()
	docker := NewDocker(*config)
	buildImage(docker)
}

func buildImage(docker Docker) {
	docker.Login()
	docker.Build()
	docker.PushTags()
}
