package main

import (
    "log"
    "flag"
    "net/http"
    "code.google.com/p/go.net/websocket"
)

type Page struct {
    Template string
    Content interface{}
}

type DummyPage struct {
    Text string
}

var addr = flag.String("addr", ":8080", "http service address")

func main() {
    flag.Parse()
    go h.run()
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/client", chatHandler)
    http.Handle("/wsclient", websocket.Handler(wsHandlerClient))
    http.Handle("/wsdaemon", websocket.Handler(wsHandlerDaemon))
    if err := http.ListenAndServe(*addr, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
