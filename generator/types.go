package generator

import ()

type Project struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Url         string `yaml:"url"`
}

type Experience struct {
	Title       string `yaml:"title"`
	StartDate   string `yaml:"startDate"`
	Description string `yaml:"description"`
}

type Profile struct {
	Name    string `yaml:"name"`
	Tagline string `yaml:"tagline"`
}
