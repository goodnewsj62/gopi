package main

import (
	"net/http"
)


func main(){
	// http.ListenAndServe(":8010",nil)

	server := http.Server{
		Addr: "0.0.0.0:8010",
		Handler: nil,
	}

	server.ListenAndServe()
}