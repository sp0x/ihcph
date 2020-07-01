package main

import (
	"github.com/gorilla/mux"
	"github.com/sp0x/ihcph/function"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	log.Printf("server started on: %s", ":8089")

	//r.Use(VerifyHTTPRequest)
	r.HandleFunc("/function/method", function.TestRequest)
	s := &http.Server{
		Addr:    ":8089",
		Handler: r,
	}
	_ = s.ListenAndServe()
	//log.Fatal(s.ListenAndServe())
}
