package models

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var cyoaTemplate = makeTemplate("../.././html/adventure.html")

func makeTemplate(filename string) *template.Template {
	return template.Must(template.ParseFiles(filename))
}

// StoryArc is a path in a story
type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// Story is an adventure with story arcs
type Story map[string]StoryArc

// ExplainStoryArc
func (s StoryArc) ExplainStoryArc() string {
	return fmt.Sprintf("%s\n\n\t%s\n", s.Title, strings.Join(s.Story, "\n\t"))
}

// TraverseArcs
func (story Story) TraverseArcs(arc string, action func(StoryArc) string) {
	chosenArc := story[arc]
	if nextArc := action(chosenArc); nextArc != "" {
		story.TraverseArcs(nextArc, action)
	}
}

// httpHandler
func (story Story) httpHandler(arc string, w http.ResponseWriter) error {
	t, err := cyoaTemplate.Clone()
	if err != nil {
		return err
	}
	data := story[arc]
	// log.Printf("%+v", data)
	return t.Execute(w, data)
}

// ServeHTTP
func (story Story) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc := strings.TrimSpace(r.URL.Path)
	if arc == "/" {
		arc = "/intro"
	}
	arc = strings.TrimSuffix(arc, "/")

	if err := story.httpHandler(arc[1:], w); err != nil {
		log.Println("error rendering template", err)
	}
}
