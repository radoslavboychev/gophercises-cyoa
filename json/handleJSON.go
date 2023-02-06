package json

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"text/template"

	"github.com/radoslavboychev/gophercises-cyoa/models"
)

// unmarshals
func ReadJSON(filename string) (models.Story, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	var data map[string]models.Adventure
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return nil, errors.New("failed to unmarshal")
	}

	return data, nil
}

func MakeTemplate(filename string) *template.Template {
	return template.Must(template.ParseFiles(filename))
}

