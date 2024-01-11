package main

import (
	"log"
	"net/http"
	"text/template"
)

type Message struct {
	Sender  string
	Content string
}

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/index.html"))

		messages := map[string][]Message{
			"Messages": {
				{Sender: "Anya", Content: "Hello! This is Anya 😺"},
				{Sender: "Self", Content: "Hi!"},
				{Sender: "Anya", Content: "How are you?"},
			},
		}

		tmpl.Execute(w, messages)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		sender := "Self"
		content := r.PostFormValue("content")
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.ExecuteTemplate(w, "message-list-element", Message{Sender: sender, Content: content})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-message/", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
