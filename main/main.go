package main

import (
	"fmt"
	"net/http"

	"github.com/Sumedhvats/url_shortner/urlHandler"
)

func Main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/github": "https://github.com/sumedhvats/",
		"youtube": "https://www.youtube.com/",
	}
	mapHandler := urlHandler.MapHandler(pathsToUrls, mux)
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlHandler.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", nil)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "job dedo")
}
