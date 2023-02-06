package models

import (
	"log"
	"net/http"
	"strings"
)

type Data struct {
	Adventures []Adventure
}

// Adventure represents the entire adventure
type Adventure struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Story map[string]Adventure

func (s Story) httpHandler(arc string, w http.ResponseWriter) {
	t, err := cyoaTemplate.Clone()
}

func (s Story) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc := strings.TrimSpace(r.URL.Path)

	if arc == "/" {
		arc = "/intro"
	}

	if strings.HasSuffix(arc, "/") {
		arc = strings.TrimSuffix(arc, "/")
	}

	if err := s.httpHandler(arc[1:], w); err != nil {
		log.Println("error rendering template", err)
	}
}
