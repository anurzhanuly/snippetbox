package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)

		return
	}

	_, err := w.Write([]byte("Hello from Snippetbox"))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)

		return
	}

	_, err = fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	if err != nil {
		return
	}
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method not allowed", 405)

		return
	}

	_, err := w.Write([]byte("createSnippet"))
	if err != nil {
		log.Fatal(err.Error())
	}
}