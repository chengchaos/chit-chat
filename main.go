package main

import (
	"net/http"
	"time"
)

func main() {
	p("Chitchat", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handle functions defined in other files
	//

	// index
	//mux.HandleFunc("/", index)
	//// error
	//mux.HandleFunc("/err", err)
	//
	//// defined in route_auth.go
	//mux.HandleFunc("/log-in", login)
	//mux.HandleFunc("/log-out", loutout)
	//mux.HandleFunc("/sign-up", signup)
	//mux.HandleFunc("/sign-up-account", singupAccount)
	//mux.HandleFunc("/authenticate", authenticate)
	//
	//// defined in route_thread.go
	//mux.HandleFunc("/thread/new", newThread)
	//mux.HandleFunc("/thread/create", createThread)
	//mux.HandleFunc("/thread/post", postThread)
	//mux.HandleFunc("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
