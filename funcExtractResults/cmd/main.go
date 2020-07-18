package main

import (
	"github.com/gorilla/mux"
	"github.com/sp0x/ihcph"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/extract", ihcph.ExtractResults)
	s := &http.Server{
		Addr:    ":8089",
		Handler: r,
	}
	log.Printf("server starting on: %s", ":8089")
	_ = s.ListenAndServe()
}
