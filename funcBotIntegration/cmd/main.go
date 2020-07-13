package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/sp0x/ihcph"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/bots/new", ihcph.NewBotIntegration)
	s := &http.Server{
		Addr:    ":8089",
		Handler: r,
	}
	log.Printf("server starting on: %s", ":8089")
	_ = s.ListenAndServe()
}
