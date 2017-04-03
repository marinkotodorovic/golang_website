package main

import (
	"html/template"
	"log"
	"net/http"
)

type pageData struct {
	Title     string
	FirstName string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.ListenAndServe(":8080", nil)
}

func sendView(w http.ResponseWriter, pd pageData) {
	w.Header().Set("Content-Type", "text/html")
	err := tpl.ExecuteTemplate(w, "templist.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	pd := pageData{Title: "Index"}
	sendView(w, pd)
}

func about(w http.ResponseWriter, req *http.Request) {
	pd := pageData{Title: "About"}
	sendView(w, pd)
}

func contact(w http.ResponseWriter, req *http.Request) {
	pd := pageData{Title: "Contact"}
	sendView(w, pd)
}

func apply(w http.ResponseWriter, req *http.Request) {
	pd := pageData{Title: "Apply"}
	w.Header().Set("Content-Type", "text/html")
	if req.Method == http.MethodPost {
		pd.FirstName = req.FormValue("fname")
	}
	err := tpl.ExecuteTemplate(w, "createview.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
