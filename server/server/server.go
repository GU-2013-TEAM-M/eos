package main

import (
    "fmt"
    "net/http"
    "eos/server/db"
)

type Page struct {
    Template, Text string
}

// FIXME: change with an appropriate handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]
    if path == "" {
        page := Page{ Template: "index.html", Text: "Hello World!" }
        ServeHtml(w, &page)
    } else if path == "db" {
        db.Connect()
        fmt.Fprintf(w, db.Test())
    } else {
        fmt.Fprintf(w, "Hi there, you've asked for %s", path)
    }
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.ListenAndServe(":8080", nil)
}
