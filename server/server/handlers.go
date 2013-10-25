package main

import (
    "fmt"
    "net/http"
    "eos/server/db"
    "code.google.com/p/go.net/websocket"
)

//-------------------------------------------------------
// WebSockets handlers
//-------------------------------------------------------
// a WebSocket handler for dealing with clients
func wsHandlerClient(ws *websocket.Conn) {
    u := &User{SessionId: "dummy"}
    c := &Connection{send: make(chan string, 256), ws: ws, owner: u}
    u.c = c // backlink for authorisation
    h.register <- c
    defer func() { h.unregister <- c }()
    go c.writer()
    c.reader()
}

// a WebSocket handler for dealing with daemons
func wsHandlerDaemon(ws *websocket.Conn) {
    d := &Daemon{Id: "dummy", IP: "0.0.0.0"}
    c := &Connection{send: make(chan string, 256), ws: ws, owner: d}
    d.c = c // backlink for authorisation
    h.register <- c
    defer func() { h.unregister <- c }()
    go c.writer()
    c.reader()
}

//-------------------------------------------------------------
// nonsense handlers, to be deleted
// FIXME: change with an appropriate handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    if path == "" {
        page := Page{ Template: "index.html", Content: DummyPage{ Text: "Hello World!" } }
        ServeHtml(w, &page)
    } else if path == "db" {
        db.Connect()
        fmt.Fprintf(w, db.Test())
    } else {
        fmt.Fprintf(w, "Hi there, you've asked for %s", path)
    }
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
    page := Page{ Template: "client.html", Content: r.Host }
    ServeHtml(w, &page)
}

