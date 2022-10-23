package main

import (
	"net/http"
	"time"

	"github.com/chengchaos/chit-chat/utils"

	"github.com/chengchaos/chit-chat/ctrl"
)

func main() {
	config := utils.Config
	utils.P("Chitchat", utils.Version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handle functions defined in other files
	//

	// index
	mux.HandleFunc("/", ctrl.Index)
	mux.HandleFunc("/index", ctrl.Index)
	// error
	mux.HandleFunc("/err", ctrl.Err)

	// defined in route_auth.go
	mux.HandleFunc("/login", ctrl.Login)
	mux.HandleFunc("/signup", ctrl.Signup)
	mux.HandleFunc("/signupaccount", ctrl.SignupAccount)
	mux.HandleFunc("/logout", ctrl.Logout)
	mux.HandleFunc("/authenticate", ctrl.Authenticate)

	//// defined in route_thread.go
	mux.HandleFunc("/thread/new", ctrl.NewThread)
	mux.HandleFunc("/thread/create", ctrl.CreateThread)
	mux.HandleFunc("/thread/post", ctrl.PostThread)
	mux.HandleFunc("/thread/read", ctrl.ReadThread)

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
