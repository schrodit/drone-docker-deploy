package main

import (
	"fmt"
	"log"
	"os"
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
	tags := NewTags()
	envVars := NewEnvVars(tags)
	run(envVars)
	fmt.Println("succesfully published images")
}

func debugGit() {
	cmd := exec.Command("env")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("error in debug\n%v", err)
	}

	cmd = exec.Command("cat", "/root/.netrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("error in debug\n%v", err)
	}
}

func run(e EnvVars) {
	//create Config out of Environment variables
	config := e.Get()
	debugGit()
	docker := NewDocker(*config)
	buildImage(docker)
}

func buildImage(docker Docker) {
	docker.Login()
	docker.Build()
	docker.PushTags()
}
