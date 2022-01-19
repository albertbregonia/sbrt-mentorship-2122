package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle(`/`, http.FileServer(http.Dir(`frontend`)))
	log.Println(`Server Started`)
	log.Fatal(http.ListenAndServe(`:54000`, nil))
}
