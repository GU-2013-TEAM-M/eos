package main

import (
    "net/http"
    "html/template"
)

func ServeHtml(w http.ResponseWriter, page *Page) {
    t, _ := template.ParseFiles("templates/" + page.Template)
    t.Execute(w, page.Content)
}
