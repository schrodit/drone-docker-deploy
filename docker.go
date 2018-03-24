package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Docker interface {
	Login()
	Build()
	Push(string)
	Tag(string, string)
	PushTags()
}

type docker struct {
	Config
}

func NewDocker(c Config) Docker {
	d := docker{c}
	return &d
}

func (d *docker) Login() {
	if d.Config.Username != "" {
		err := execDocker("login", "-u", d.Username, "-p", d.Password, d.Registry)
		if err != nil {
			log.Fatal("login failed")
			os.Exit(1)
		}
	}
}

func (d *docker) Build() {
	err := execDocker("build", "-t", d.Image, "-f", d.Dockerfile, d.Dir)
	if err != nil {
		log.Fatalf("failed to build image %s in directory %s", d.Image, d.Dir)
		os.Exit(1)
	}
}

func (d *docker) Push(image string) {
	err := execDocker("push", image)
	if err != nil {
		log.Fatalf("failed to push image: %v", image)
		os.Exit(1)
	}
}

func (d *docker) Tag(source, target string) {
	err := execDocker("tag", source, target)
	if err != nil {
		log.Fatalf("failed to tag image: %v", source)
		os.Exit(1)
	}
}

func (d *docker) PushTags() {
	for _, tag := range d.Config.Tags {
		image := fmt.Sprintf("%s:%s", d.Config.Image, tag)
		d.Tag(d.Config.Image, image)
		d.Push(image)
	}
}

func execDocker(args ...string) error {
	return execCommand("docker", args...)
}

func execCommand(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
