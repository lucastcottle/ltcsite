package handlers

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type AboutMe struct {
	name      string `yaml:"name"`
	tagline   string `yaml:"tagline"`
	githubUrl string `yaml:"githubUrl"`
}

func ParseYaml(filename string) {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading yaml file: %v", err)
	}

	hp := AboutMe{}

	err = yaml.Unmarshal(yamlFile, &hp)
	if err != nil {
		log.Fatalf("error parsing yaml: %v", err)
	}
}
