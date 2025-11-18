package handlers

import (
	"net/http"
	"text/template"

	"github.com/lucastcottle/ltcsite/generator"
)

func ProfileHandler(t *template.Template, p generator.Profile) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "index", p)
	}
}

func ProjectsHandler(t *template.Template, p []generator.Project) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "projects", p)
	}
}

func ExperienceHandler(t *template.Template, p []generator.Experience) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "experience", p)
	}
}
