package generator

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadProjects(path string) ([]Project, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var out []Project
	err = yaml.Unmarshal(b, &out)
	return out, err
}

func LoadExperience(path string) ([]Experience, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var out []Experience
	err = yaml.Unmarshal(b, &out)
	return out, err
}

func LoadProfile(path string) (Profile, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Profile{}, err
	}

	var out Profile
	err = yaml.Unmarshal(b, &out)
	return out, err
}

func MustLoadProfile() Profile {
	p, err := LoadProfile("internal/data/profile.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded profile: %+v\n", p)
	return p
}

func MustLoadExperience() []Experience {
	e, err := LoadExperience("internal/data/experience.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return e
}

func MustLoadProjects() []Project {
	p, err := LoadProjects("internal/data/projects.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return p
}
