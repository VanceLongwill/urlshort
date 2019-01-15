package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	".."
)

func main() {

	yamlFilename := flag.String("f", "mock.yaml", "yaml file")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	file, err := ioutil.ReadFile(*yamlFilename)
	if err != nil {
		fmt.Print(err, "\nFailed to read the YAML file provided")
		return
	}

	yamlHandler, yamlErr := urlshort.YAMLHandler([]byte(file), mapHandler)

	if yamlErr != nil {
		panic(yamlErr)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
