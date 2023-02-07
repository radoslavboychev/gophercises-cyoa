package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/radoslavboychev/gophercises-cyoa/models"
)

func cliHandler(s models.StoryArc) string {
	fmt.Fprintln(os.Stdout, s.ExplainStoryArc())
	if len(s.Options) > 0 {
		fmt.Fprintln(os.Stdout, "Choose an option:")
		for index, option := range s.Options {
			fmt.Fprintf(os.Stdout, "\n%d %s: %s\n", index+1, option.Arc, option.Text)
		}
		var choice int
		if _, err := fmt.Fscanf(os.Stdin, "%d", &choice); err != nil {
			if err == io.EOF {
				return ""
			}
			log.Fatalln("Error reading option:", err)
		}
		fmt.Fprintln(os.Stdout, choice, s.Options[0].Arc)
		if nextArc := s.Options[choice-1].Arc; nextArc != "" {
			return nextArc
		}
	}
	return ""
}

// parse a story from a map (loaded from a JSON file) into a story object
func parseStory(storyMap []byte) (story models.Story) {
	if err := json.Unmarshal(storyMap, &story); err != nil {
		log.Fatalf("Error parsing story map: %s", err)
	}
	return
}
