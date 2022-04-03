package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
)

func main() {
	http.Handle(`/`, http.FileServer(http.Dir(`frontend`)))
	http.HandleFunc(`/login`, LoginHandler) //handle login
	// http.HandleFunc(`/chat`, ChatHandler)   //handle chat
	log.Println(`Server Initialized`)
	log.Fatal(http.ListenAndServeTLS(`:54000`, `server.crt`, `server.key`, nil)) //null, undefined, NULL
}

var (
	sessions   = make(map[string]int)
	sessionMtx = sync.RWMutex{}
)

//user database
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if e := r.ParseForm(); e != nil { //if we parse the form and we get an error, tell the user
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	name := strings.TrimSpace(r.FormValue(`username`)) //remove excess whitespace on the username
	if name == `` || GetSession(name) > 0 {            //if the username is blank or taken, tell the user its invalid
		http.Error(w, `error: invalid username`, http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, AddSession(name)) //send the session id to the user
}

func GetSession(username string) int {
	sessionMtx.RLock()         //lock it so anybody can read but no changing
	defer sessionMtx.RUnlock() //once we return unlock
	return sessions[username]  //return the sessionID
}

func AddSession(username string) int {
	sessionMtx.RLock()         //lock it so anybody can read but no changing
	defer sessionMtx.RUnlock() //once we return unlock
	sessions[username] = rand.Intn(1_000_000) + 1
	return sessions[username]
}
