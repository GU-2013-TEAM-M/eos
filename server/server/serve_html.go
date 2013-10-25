package main

import (
    "net/http"
    "html/template"
)

//-------------------------------------------------------
// Data Structures
//-------------------------------------------------------
type Page struct {
    Template string
    Content interface{}
}

type DummyPage struct {
    Text string
}

//-------------------------------------------------------
// Functions to serve html templates
//-------------------------------------------------------
func ServeHtml(w http.ResponseWriter, page *Page) {
    t, _ := template.ParseFiles("static/templates/" + page.Template)
    t.Execute(w, page.Content)
}
