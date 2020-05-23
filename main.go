package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	data := map[string]string{
		"Title": greeting("Code.education Rocks!"),
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}

func greeting(text string) string {
	if (len(text) == 0 ) {
		return "empty :("
	} else {
		return text
	}
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
