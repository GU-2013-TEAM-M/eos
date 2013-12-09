package main

import (
    "log"
    "flag"
    "net/http"
    "code.google.com/p/go.net/websocket"
    "eos/server/db"
)

//-------------------------------------------------------
// variables
//-------------------------------------------------------
var addr = flag.String("addr", ":8080", "http service address")
var TEST = false

//-------------------------------------------------------
// main execution block
//-------------------------------------------------------
func main() {
    db.Connect()
    flag.Parse()
    RegisterAllHandlers()
    go h.run()

    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/client", chatHandler)
    http.Handle("/static/", http.FileServer(http.Dir("")))
    http.Handle("/wsclient", websocket.Handler(wsHandlerClient))
    http.Handle("/wsdaemon", websocket.Handler(wsHandlerDaemon))
    if err := http.ListenAndServe(*addr, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
    NewOrg("Anonymous")
}
