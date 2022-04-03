package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle(`/`, http.FileServer(http.Dir(`frontend`))) //root of the page is the frontend
	log.Println(`Server Started`)                           //print that the server started with a time stamp
	log.Fatal(http.ListenAndServe(`:54000`, nil))           //print out the error and exit the program if starting the server fails
}
