package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page requested")
	RenderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("about page requested")
	RenderTemplate(w, "about.html")
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedtemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedtemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template")
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil)
}
