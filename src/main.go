package main

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Registry string
	Image    string
	Dir      string
	Username string
	Password string
	Tags     []string
	JobNum   string
}

func main() {
	config := getEnvVars()

	docker := NewDocker(*config)
	docker.Login()
	docker.Build()

	for _, tag := range config.Tags {
		newTag := fmt.Sprintf("%s-%s", tag, config.JobNum)
		image := fmt.Sprintf("%s:%s", config.Image, newTag)
		docker.Tag(config.Image, image)
		docker.Push(image)
	}
}

func getEnvVars() *Config {
	config := Config{}

	config.Registry = os.Getenv("PLUGIN_REGISTRY")

	config.Image = os.Getenv("PLUGIN_IMAGE")
	if config.Image == "" {
		fmt.Errorf("parameter image is required")
		os.Exit(1)
	}
	config.Dir = os.Getenv("PLUGIN_DIRECTORY")
	if config.Registry == "" {
		config.Dir = "."
	}

	config.JobNum = os.Getenv("DRONE_JOB_NUMBER")

	//get credentials
	secret := os.Getenv("PLUGIN_SECRETS")
	if secret != "" {
		secrets := strings.Split(secret, ",")
		config.Username = secrets[0]
		config.Password = secrets[1]
	}

	return &config
}
