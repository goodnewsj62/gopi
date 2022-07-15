package main

import (
	"fmt"
	"net/http"
)


type MyHadler struct{}

func (f *MyHadler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Hello world again", "\n",r,"url:",r.URL)
}

func main(){
	
	server := &http.Server{
		Addr: "0.0.0.0:8010",
		Handler: &MyHadler{},
	}

	server.ListenAndServe()
}	