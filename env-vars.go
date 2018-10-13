package main

import (
	"log"
	"os"
)

type EnvVars interface {
	Get() *Config
}

type envVars struct {
	tags Tags
}

func NewEnvVars(t Tags) EnvVars {
	return &envVars{t}
}

func (e *envVars) Get() *Config {
	config := Config{}

	config.Registry = os.Getenv("PLUGIN_REGISTRY")
	config.Image = os.Getenv("PLUGIN_REPO")
	if config.Image == "" {
		log.Fatal("parameter 'image' is required")
		os.Exit(1)
	}

	config.Dockerfile = os.Getenv("PLUGIN_DOCKERFILE")
	config.Dir = GetEnv("PLUGIN_DIRECTORY", ".")
	config.ImageTagsFile = GetEnv("PLUGIN_IMAGETAGSFILE", ".image_tags")
	config.BuildEvent = GetEnv("DRONE_BUILD_EVENT", "")

	if os.Getenv("PLUGIN_ADDJOBNUMBER") == "true" {
		config.AddJobNumber = true
	} else {
		config.AddJobNumber = false
	}

	if os.Getenv("PLUGIN_LATEST") == "true" {
		config.Latest = true
	} else {
		config.Latest = false
	}

	if os.Getenv("PLUGIN_USEGITTAG") == "true" {
		config.UseGitTag = true
		config.GitTag = os.Getenv("DRONE_TAG")
		if config.GitTag == "" {
			config.GitTag = e.tags.GetNewestGitTag()
			if config.GitTag == "" {
				config.UseGitTag = false
				log.Println("cannot get git tag, use .tags file")
			}
		}
	} else {
		config.UseGitTag = false
	}

	config.JobNum = os.Getenv("DRONE_BUILD_NUMBER")

	//get credentials
	config.Username = os.Getenv("DOCKER_USERNAME")
	config.Password = os.Getenv("DOCKER_PASSWORD")

	//get tags
	config.Tags = e.tags.GetTags(config)

	return &config
}

func GetEnv(envName, defaultValue string) string {
	env := os.Getenv(envName)
	if env == "" {
		env = defaultValue
	}
	return env
}
