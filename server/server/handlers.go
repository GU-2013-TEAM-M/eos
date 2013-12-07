package main

import (
    "fmt"
    "time"
    "net/http"
    "eos/server/db"
    "code.google.com/p/go.net/websocket"
)

//-------------------------------------------------------
// WebSockets handlers
//-------------------------------------------------------
// a WebSocket handler for dealing with clients
func wsHandlerClient(ws *websocket.Conn) {
    id := fmt.Sprintf("%d", time.Now().Nanosecond())
    u := &User{Id: id, SessionId: "dummy", OrgId: NO_ORG}
    c := &Connection{send: make(chan string, 256), ws: ws, owner: u}
    u.c = c // backlink for authorisation
    h.register <- c
    defer func() { h.unregister <- c }()
    go c.writer()
    c.reader()
}

// a WebSocket handler for dealing with daemons
func wsHandlerDaemon(ws *websocket.Conn) {
    id := fmt.Sprintf("%d", time.Now().Nanosecond())
    d := &Daemon{Id: id, IP: "0.0.0.0", OrgId: NO_ORG}
    c := &Connection{send: make(chan string, 256), ws: ws, owner: d}
    d.c = c // backlink for authorisation
    h.register <- c
    defer func() { h.unregister <- c }()
    go c.writer()
    c.reader()
}

//-------------------------------------------------------
// Message handlers
//-------------------------------------------------------
func RegisterAllHandlers() {
    RegisterHandler("loginCheck", LoginCheckHandler)
    RegisterHandler("login", LoginHandler)
    RegisterHandler("logout", LogoutHandler)
    RegisterHandler("daemons", DaemonsHandler)
    RegisterHandler("daemon", DaemonHandler)
    RegisterHandler("control", ControlHandler)
}

//-------------------------------------------------------------
// nonsense handlers, to be deleted
// FIXME: change with an appropriate handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    if path == "" {
        http.ServeFile(w, r, "./static/templates/index.html")
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

