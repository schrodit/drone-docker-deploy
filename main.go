package main

import (
	"fmt"
	"log"
	"os/exec"
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
	setupNetrc()
	tags := NewTags()
	envVars := NewEnvVars(tags)
	run(envVars)
	fmt.Println("succesfully published images")
}

func setupNetrc() {
	cmd := exec.Command("/setup/Netrc.sh")
	if err := cmd.Run(); err != nil {
		log.Printf("cannot setup netrc\n%v", err)
		return
	}
	log.Println("succesfully setup .netrc file")
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
