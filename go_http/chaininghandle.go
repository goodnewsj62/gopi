package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world again")
}


func log(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("\nlog was called")
		h.ServeHTTP(w,r)
	})
}


func protect(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Do some protect stuff")
		h.ServeHTTP(w,r)
	})
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8010",
	}
	myhandler := MyHandler{}
	http.Handle("/", protect(log(&myhandler)))
	server.ListenAndServe()
}