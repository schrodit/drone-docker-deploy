package main

import (
	"fmt"
	"log"
	"os"
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
	fmt.Println("starting deployment")
	config := getEnvVars()

	docker := NewDocker(*config)
	if config.Username != "" {
		docker.Login()
	}
	docker.Build()

	for _, tag := range config.Tags {
		newTag := fmt.Sprintf("%s-%s", tag, config.JobNum)
		image := fmt.Sprintf("%s:%s", config.Image, newTag)
		docker.Tag(config.Image, image)
		docker.Push(image)
	}
	fmt.Println("succesfully published images")
}

func getEnvVars() *Config {
	config := Config{}

	config.Registry = os.Getenv("PLUGIN_REGISTRY")

	config.Image = os.Getenv("PLUGIN_REPO")
	if config.Image == "" {
		log.Fatal("parameter 'image' is required")
		os.Exit(1)
	}
	config.Dir = os.Getenv("PLUGIN_DIRECTORY")
	if config.Dir == "" {
		config.Dir = "."
	}

	config.JobNum = os.Getenv("DRONE_BUILD_NUMBER")

	//get credentials
	config.Username = os.Getenv("DOCKER_USERNAME")
	config.Password = os.Getenv("DOCKER_PASSWORD")

	//get tags
	var err error
	config.Tags, err = ReadTagsFile(".tags")
	if err != nil || len(config.Tags) == 0 {
		config.Tags = []string{"latest"}
	}

	return &config
}
