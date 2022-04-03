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

type SafeSocket struct {
	*websocket.Conn
	sync.RWMutex
}

func (ws *SafeSocket) SendMessage(msg Message) error {
	ws.Lock()
	defer ws.Unlock()
	return ws.WriteJSON(msg)
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  512,
		WriteBufferSize: 512,
	}
)

func main() {
	http.Handle(`/`, http.FileServer(http.Dir(`frontend`)))
	http.HandleFunc(`/login`, LoginHandler)
	http.HandleFunc(`/chat`, ChatHandler)
	go Broadcast()
	log.Println(`Server Initialized`)
	log.Fatal(http.ListenAndServeTLS(`:54000`, `server.crt`, `server.key`, nil))
}

//user database

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if e := r.ParseForm(); e != nil { //parse the form, if it fails let the user know
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	name := strings.TrimSpace(r.FormValue(`username`)) //remove excess whitespace with trim
	if name == `` || GetSession(name) > 0 {            //if the username is blank or someone with that username exists
		http.Error(w, `error: invalid username`, http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, AddSession(name)) //send the session id to the user
}

var (
	users       = make(map[string]*SafeSocket)
	usersMtx    = sync.RWMutex{}
	sessions    = make(map[string]int)
	sessionsMtx = sync.RWMutex{}
)

func GetSession(username string) int {
	sessionsMtx.RLock()
	defer sessionsMtx.RUnlock()
	return sessions[username]
}

func AddSession(username string) int { //what is the problem with this setup?
	sessionsMtx.Lock()
	defer sessionsMtx.Unlock()
	sessions[username] = rand.Intn(1_000_000) + 1 //can never be 0 as that will be the `not found` case
	return sessions[username]
}

func RemoveSession(username string) {
	sessionsMtx.Lock()
	defer sessionsMtx.Unlock()
	delete(sessions, username)
}

func AddUser(name string, ws *SafeSocket) {
	usersMtx.Lock()
	defer usersMtx.Unlock()
	users[name] = ws
}

func RemoveUser(name string) {
	usersMtx.Lock()
	defer usersMtx.Unlock()
	delete(users, name)
}

//chat functions

var broadcast = make(chan Message)

type Message struct {
	Content string `json:"content"`
	Sender  string `json:"sender"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	if e := r.ParseForm(); e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	name := strings.TrimSpace(r.FormValue(`username`))
	sessionID, e := strconv.Atoi(strings.TrimSpace(r.FormValue(`sessionID`)))
	if e != nil || name == `` { //if we have any invalid parameters
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	if GetSession(name) != sessionID { //if the saved sessionID does not match the given sessionID then break
		http.Error(w, e.Error(), http.StatusForbidden)
		return
	}
	ws, e := wsUpgrader.Upgrade(w, r, nil)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	safeWS := &SafeSocket{ws, sync.RWMutex{}}
	AddUser(name, safeWS)
	defer RemoveUser(name)
	defer RemoveSession(name)
	var msg Message
	for {
		if e := safeWS.ReadJSON(&msg); e != nil { //infinitely read in messages
			log.Println(e)
			return
		}
		msg.Sender = name
		broadcast <- msg //infinitely broadcast them out
	}
}

func Broadcast() {
	for msg := range broadcast {
		usersMtx.Lock()
		for _, user := range users { //send everyone that broadcasted message
			user.SendMessage(msg)
		}
		usersMtx.Unlock()
	}
}
