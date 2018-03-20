package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Docker interface {
	Login()
	Build()
	Push(string)
	Tag(string, string)
}

type docker struct {
	Config
}

func NewDocker(c Config) Docker {
	d := docker{c}
	return &d
}

func (d *docker) Login() {
	_, err := execDocker("login", "-u", d.Username, "-p", d.Password, d.Registry)
	if err != nil {
		fmt.Errorf("login failed")
		os.Exit(1)
	}
}

func (d *docker) Build() {
	_, err := execDocker("build", "-t", d.Image, d.Dir)
	if err != nil {
		fmt.Errorf("failed to build image")
		os.Exit(1)
	}
}

func (d *docker) Push(image string) {
	_, err := execDocker("push", image)
	if err != nil {
		fmt.Errorf("failed to push image: %v", image)
		os.Exit(1)
	}
}

func (d *docker) Tag(source, target string) {
	_, err := execDocker("tag", source, target)
	if err != nil {
		fmt.Errorf("failed to tag image: %v", source)
		os.Exit(1)
	}
}

func execDocker(args ...string) ([]byte, error) {
	return execCommand("docker", args...)
}

func execCommand(cmd string, args ...string) ([]byte, error) {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	return c.Output()
}
