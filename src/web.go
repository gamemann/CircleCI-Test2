package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var tem *template.Template

func handler(w http.ResponseWriter, r *http.Request) {
	// Execute template.
	tem.Execute(w, nil)
}

func createwebserver(port int) error {
	// Create new router.
	r := mux.NewRouter()

	// Handle routes.
	r.HandleFunc("/", handler)

	// Serve static files.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// Create main route.
	r.Handle("/", r)

	// We want to create the server itself so we can specify the read and write timeouts.
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      r,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	// Parse template.
	var err error
	tem, err = template.ParseFiles("templates/index.html")

	if err != nil {
		return err
	}

	// Initiate the server and have it listen on whatever port we specified (8808).
	err = srv.ListenAndServe()

	return err
}

func main() {
	// Create the web server and if it returns false, log the error.
	log.Fatal(createwebserver(8808))
}
