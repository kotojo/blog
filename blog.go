package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type BlogPost struct {
	Title string
	Body  []byte
}

func loadBlogPost(title string) (*BlogPost, error) {
	filename := "views/blogs/" + title + ".md"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &BlogPost{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, bp *BlogPost) {
	t, _ := template.ParseFiles("views/" + tmpl + ".html")
	err := t.Execute(w, bp)
	if err != nil {
		fmt.Fprint(w, "Wow! Something went really wrong over here!")
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(("/posts/")):]
	bp, err := loadBlogPost(title)
	if err != nil {
		bp = &BlogPost{Title: title, Body: []byte("Looks like I never wrote this post!")}
	}
	renderTemplate(w, "blog", bp)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("views/index.html")
	if err != nil {
		fmt.Fprintf(w, "Sorry! Couldn't find that page!")
	}
	fmt.Fprint(w, string(file))
}

func main() {
	http.HandleFunc("/posts/", postHandler)
	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
