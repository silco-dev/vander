package main

import "github.com/gorilla/mux"

func (a *API) setRouter() {
	a.Router = mux.NewRouter()

	a.Router.Use(a.Authenticate)

	a.Router.HandleFunc("/", a.apiRoot)
}
