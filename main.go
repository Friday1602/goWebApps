package main

import (
	"log"
	"net/http"
)

// home handler func
func home(w http.ResponseWriter, r *http.Request) {
	// protect the servemux to treat this like catch all pattern
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox"))
}

// snippet view handler func
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// create snippet handler func
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// router (servemux) mapping url pattern and corresponding handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snipper/create", snippetCreate)
	// establish a web server and listen for incoming requests
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
