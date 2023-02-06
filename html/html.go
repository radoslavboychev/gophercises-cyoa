package html

import (
	"html/template"
	"log"
)

func LoadTemplate(filename string) {
	t := template.New("Adventure")
	t, err := t.ParseFiles(filename)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
}

func ParseTemplate(filename string) (*template.Template, error) {
	parsedTemplate, err := template.ParseFiles(filename)
	if err != nil {
		return nil, err
	}
	return parsedTemplate, nil
}
