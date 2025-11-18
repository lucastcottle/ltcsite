package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/lucastcottle/ltcsite/generator"
	"github.com/lucastcottle/ltcsite/handlers"
)

func main() {
	t := template.Must(template.ParseGlob("templates/*.html"))

	profile := generator.MustLoadProfile()
	projects := generator.MustLoadProjects()
	experience := generator.MustLoadExperience()

	http.Handle("/", handlers.ProfileHandler(t, profile))
	http.Handle("/projects", handlers.ProjectsHandler(t, projects))
	http.Handle("/experience", handlers.ExperienceHandler(t, experience))

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
