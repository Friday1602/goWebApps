package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// create snippet handler func
// add POST only method to snippetCreate handler
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// Add allow: POST to the header map
		w.Header().Set("Allow", http.MethodPost) // use MethodPost instead of "POST"
		// must call writeheader before write to send non-200 status code
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// use this http.Error instead
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // use statusMethodNotAllowed instead of 405
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// router (servemux) mapping url pattern and corresponding handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	// establish a web server and listen for incoming requests
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
