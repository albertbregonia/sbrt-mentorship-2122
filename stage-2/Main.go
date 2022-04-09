package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

func main() {
	http.Handle(`/`, http.FileServer(http.Dir(`frontend`)))
	http.HandleFunc(`/login`, LoginHandler) //handle login
	http.HandleFunc(`/chat`, ChatHandler)   //handle chat
	go Broadcast()
	log.Println(`Server Initialized`)
	log.Fatal(http.ListenAndServeTLS(`:443`, `server.crt`, `server.key`, nil)) //null, undefined, NULL
}

var (
	sessions   = make(map[string]*User) //database
	sessionMtx = sync.RWMutex{}
)

type User struct {
	sessionID int
	ws        SafeSocket
}

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
	if sessions[username] == nil {
		return 0
	}
	return sessions[username].sessionID //return the sessionID
}

func AddSession(username string) int {
	sessionMtx.Lock()         //lock it so anybody can read but no changing
	defer sessionMtx.Unlock() //once we return unlock
	sessions[username] = &User{sessionID: rand.Intn(1_000_000) + 1}
	return sessions[username].sessionID
}

func RemoveSession(username string) {
	sessionMtx.Lock()
	defer sessionMtx.Unlock()
	sessions[username].ws.Close()
	delete(sessions, username)
}

type SafeSocket struct {
	*websocket.Conn
	*sync.RWMutex
}

type Message struct {
	Content string `json:"content"`
	Sender  string `json:"sender"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  512,
		WriteBufferSize: 512,
	}
)

var broadcast = make(chan Message)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	if e := r.ParseForm(); e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	name := strings.TrimSpace(r.FormValue(`username`))
	sessionID, e := strconv.Atoi(strings.TrimSpace(r.FormValue(`sessionID`)))
	if e != nil || name == `` {
		http.Error(w, `error: invalid username or sessionID`, http.StatusBadRequest)
		return
	}
	if GetSession(name) != sessionID {
		http.Error(w, `error: sessionID does not match the given username`, http.StatusForbidden)
		return
	}
	ws, e := wsUpgrader.Upgrade(w, r, nil)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	sessionMtx.Lock()
	sessions[name].ws = SafeSocket{ws, &sync.RWMutex{}} //adds the websocket to our user
	sessionMtx.Unlock()
	defer RemoveSession(name) //once they disconnect, delete them
	var msg Message
	for {
		if e := ws.ReadJSON(&msg); e != nil {
			log.Println(e)
			return
		}
		msg.Sender = name
		broadcast <- msg
	}
}

func Broadcast() {
	for msg := range broadcast {
		sessionMtx.Lock() //prevents the sessions map from being modified before we write out to everyone
		for _, user := range sessions {
			user.ws.Lock() //prevents the read/write websocket data race
			user.ws.WriteJSON(msg)
			user.ws.Unlock()
		}
		sessionMtx.Unlock()
	}
}
