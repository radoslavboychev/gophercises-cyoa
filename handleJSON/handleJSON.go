package handlejson

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/radoslavboychev/gophercises-cyoa/models"
)

// ReadJSON loads a JSON file, reads the data from it and unmarshals it into a Story struct
func ReadJSON(filename string) (models.Story, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("failed to open file")
	}

	var data map[string]models.StoryArc
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		return nil, errors.New("failed to unmarshal")
	}

	return data, nil
}
