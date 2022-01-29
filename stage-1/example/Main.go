package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle(`/`, http.FileServer(http.Dir(`frontend`)))                    //just give the HTML/CSS/JS when loading the webserver
	log.Println(`Server Initialized`)                                          //print that the server has started
	log.Fatal(http.ListenAndServeTLS(`:443`, `server.crt`, `server.key`, nil)) //create HTTPS website
}
