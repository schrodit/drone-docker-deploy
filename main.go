package main

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	Registry   string
	Image      string
	Dockerfile string
	Dir        string
	Username   string
	Password   string
	Tags       []string
	JobNum     string
	UseGitTag  bool
	GitTag     string
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
		image := fmt.Sprintf("%s:%s", config.Image, tag)
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

	config.Dockerfile = os.Getenv("PLUGIN_DOCKERFILE")
	if config.Dockerfile == "" {
		config.Dockerfile = "Dockerfile"
	}

	config.Dir = os.Getenv("PLUGIN_DIRECTORY")
	if config.Dir == "" {
		config.Dir = "."
	}
	if os.Getenv("PLUGIN_USEGITTAG") == "true" {
		config.UseGitTag = true
		config.GitTag = os.Getenv("DRONE_TAG")
		if config.GitTag == "" {
			config.UseGitTag = false
			log.Println("cannot get git tag, use .tags file")
		}
	} else {
		config.UseGitTag = false
	}

	config.JobNum = os.Getenv("DRONE_BUILD_NUMBER")

	//get credentials
	config.Username = os.Getenv("DOCKER_USERNAME")
	config.Password = os.Getenv("DOCKER_PASSWORD")

	//get tags
	config.Tags = GetTags(config)

	return &config
}

func GetTags(config Config) []string {
	var tags []string
	if config.UseGitTag == true {
		tags = []string{config.GitTag}
	} else {
		tags, err := ReadTagsFile(".tags")
		if err != nil || len(config.Tags) == 0 {
			tags = []string{"latest"}
		}

		for i, tag := range tags {
			fmt.Println(tag)
			tags[i] = fmt.Sprintf("%s-%s", tag, config.JobNum)
		}
	}
	fmt.Println(tags)
	return tags
}
