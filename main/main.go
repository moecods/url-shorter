package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/moecods/url-shorter/urlshorter/urlshorter"
)

type URL struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}

type Response struct {
	Url string
}

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	fmt.Println(urlPath)
	urls := []URL{
		{OriginalURL: "https://www.google.com", ShortURL: "/test1"},
		{OriginalURL: "https://www.bing.com", ShortURL: "/test2"},
	}

	for _, url := range urls {
		if url.ShortURL == urlPath {
			http.Redirect(w, r, url.OriginalURL, http.StatusFound)
			return
		}
	}

	http.Error(w, "Url not found", 404)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UrlHandler)

	fmt.Println("Server is listening on port 8020...")
	http.Handle("/", mux)

	if err := http.ListenAndServe(":8020", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
