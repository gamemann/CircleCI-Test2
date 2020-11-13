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

func main() {
	var err error

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
		Addr:         ":" + strconv.Itoa(8808),
		Handler:      r,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	// Parse template.
	tem, err = template.ParseFiles("templates/index.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	// Initiate the server and have it listen on port TCP/8808.
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}
