package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/russross/blackfriday/v2"
)

type BlogPost struct {
	Title string
	Body  template.HTML
}

type FileReader func(filename string) ([]byte, error)

var templates = template.Must(template.ParseFiles("views/post.html", "view/index.html"))

func loadBlogPost(title string, fileReader FileReader) (*BlogPost, error) {
	filename := "views/posts/" + title + ".md"
	body, err := fileReader(filename)
	if err != nil {
		return nil, err
	}
	blogHtml := template.HTML(blackfriday.Run(body))
	return &BlogPost{Title: title, Body: blogHtml}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, bp *BlogPost) {
	err := templates.ExecuteTemplate(w, "views/"+tmpl+".html", bp)
	if err != nil {
		fmt.Fprint(w, err.Error(), http.StatusInternalServerError)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len(("/posts/")):]
	bp, err := loadBlogPost(title, ioutil.ReadFile)
	if err != nil {
		bp = &BlogPost{Title: title, Body: template.HTML("Looks like I never wrote this post!")}
	}
	renderTemplate(w, "post", bp)
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
