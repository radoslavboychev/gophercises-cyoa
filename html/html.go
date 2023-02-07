package html

import (
	"html/template"
	"log"
)

// LoadTemplate parses files from an HTML
func LoadTemplate(filename string) {
	t := template.New("Adventure")
	t, err := t.ParseFiles(filename)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
}

// ParseTemplate parses an HTML template from a file, returning a template
func ParseTemplate(filename string) (*template.Template, error) {
	parsedTemplate, err := template.ParseFiles(filename)
	if err != nil {
		return nil, err
	}
	return parsedTemplate, nil
}
