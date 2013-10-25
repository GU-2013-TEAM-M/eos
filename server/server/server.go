package main

import (
    "log"
    "flag"
    "net/http"
    "code.google.com/p/go.net/websocket"
)

//-------------------------------------------------------
// variables
//-------------------------------------------------------
var addr = flag.String("addr", ":8080", "http service address")

//-------------------------------------------------------
// main execution block
//-------------------------------------------------------
func main() {
    flag.Parse()
    go h.run()
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/client", chatHandler)
    http.Handle("/static", http.FileServer(http.Dir("./static/")))
    http.Handle("/wsclient", websocket.Handler(wsHandlerClient))
    http.Handle("/wsdaemon", websocket.Handler(wsHandlerDaemon))
    if err := http.ListenAndServe(*addr, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
