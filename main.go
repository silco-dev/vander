package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/silco-dev/vander/mongodb"
)

var db *mongodb.DB

const VERSION = "0.1.0"

func startAPI(quit chan bool) {

	a := &API{Connected: false, Restarting: false}
	a.setRouter()
	a.setConfig()

	serverAddr := fmt.Sprintf("%v:%v", a.Host, a.Port)
	a.Server = &http.Server{
		Addr:         serverAddr,
		Handler:      a.Router,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}
	go func() {
		if err := a.Server.ListenAndServe(); err != http.ErrServerClosed {
			log.Println("Failed to start API.", err)
		}
		quit <- true
	}()

	log.Printf("Started on http://%v/ Successfully.\n", serverAddr)
	a.Connected = true
}

func (a *API) setConfig() {
	conf := LoadConfig()
	a.Host = conf.API.Host
	a.Port = conf.API.Port
	a.Secret = conf.API.Secret

}

func main() {
	apiQuit := make(chan bool)

	config := LoadConfig()
	db = mongodb.InitDB()
	db.ConnectDB(&config.Mongo)

	go startAPI(apiQuit)

	<-apiQuit
}
