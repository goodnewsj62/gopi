package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

// func hello (w http.ResponseWriter, r *http.Request, p httprouter.Params){
// 	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
// }


type MyHandler struct{}

func (h *MyHandler) ServeHTTP (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world!")
}

func main(){
	// mux := httprouter.New()
	// mux.GET("/hello/:name", hello)
	// server := http.Server{
	// 	Addr: "127.0.0.1:8090",
	// 	Handler: mux,
	// }
	// log.Fatal( server.ListenAndServe())

	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:9091",
		Handler: &handler,
		}

		http2.ConfigureServer(&server, &http2.Server{})
		server.ListenAndServeTLS("*cert.pem","*key.pem")
}