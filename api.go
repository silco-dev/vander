package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type API struct {
	Connected             bool
	Router                *mux.Router
	Server                *http.Server
	Restarting            bool
	Secret                string
	NonSensitiveEndpoints []string
	Host                  string
	Port                  string
}

type ApiConfig struct {
	Host   string
	Port   string
	Secret string
}

type Exception struct {
	Error string `json:"error"`
}

type Response struct {
	Status string `json:"status"`
	URL    string `json:"url,omitempty"`
}

func (a *API) Error(w http.ResponseWriter, r *http.Request, code int, err error, start time.Time) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	log.Printf(
		"ERROR: %s\t%s\t%s\t",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)

}

func (a *API) Success(w http.ResponseWriter, r *http.Request, data *Response, start time.Time) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	log.Printf(
		"%s\t%s\t%s\t",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
	json.NewEncoder(w).Encode(data)
}

func (a *API) JsonSuccess(w http.ResponseWriter, r *http.Request, jsonResp []byte, start time.Time) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Write(jsonResp)
	log.Printf(
		"%s\t%s\t%s\t",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}

// Actual endpoints

func (a *API) apiRoot(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	resp := &Response{
		Status: "OK",
	}
	a.Success(w, r, resp, start)
}
