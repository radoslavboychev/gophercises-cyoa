package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/radoslavboychev/gophercises-cyoa/handlejson"
)

func main() {

	// read and load the stories from json
	story, err := handlejson.ReadJSON("../.././gopher.json")
	if err != nil {
		log.Println(err)
		return
	}

	// create the default serve mux
	r := http.DefaultServeMux

	// register the handler for the stories
	r.Handle("/", story)

	// start the server
	fmt.Println("Serving...")
	fmt.Fprintln(os.Stderr, http.ListenAndServe(":8090", story))

}
