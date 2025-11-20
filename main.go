package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucastcottle/ltcsite/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world!")
	})

	for _, value := range getYaml() {
		handlers.ParseYaml(value)
	}

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getYaml() []string {
	dirPath := "data/yaml/"

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("error reading yaml from directory '%s':\n", dirPath)
	}

	var yamlFiles []string

	for _, entry := range entries {
		if !entry.IsDir() {
			name := entry.Name()
			if strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml") {
				fullPath := filepath.Join(dirPath, name)
				yamlFiles = append(yamlFiles, fullPath)
			}
		}
	}
	return yamlFiles
}

