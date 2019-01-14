package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	".."
)

func main() {

	// yamlFilename := flag.String("f", "mock.yaml", "yaml file")
	// flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	fmt.Println(mapHandler)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	// file, err = os.Open(*yamlFilename)
	// if fileErr != nil {
	// 	fmt.Print(fileErr, "\nFailed to read the YAML file provided")
	// 	os.Exit(1)
	// }
	// yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Starting the server on :8080")
	// http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
