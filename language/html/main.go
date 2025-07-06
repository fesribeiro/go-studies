package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
	Age   int
}

func main() {
	templates = template.Must(template.ParseGlob("pages/*.html"))

	http.HandleFunc("/home", func (w http.ResponseWriter, r *http.Request) {

		user := user{
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Age:   30,
		}

		templates.ExecuteTemplate(w, "home.html", user)
	})

	fmt.Println("Server is running on http://localhost:8086")
	log.Fatal(http.ListenAndServe(":8086", nil))
}
