package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kotojo/blog/views"
	"github.com/russross/blackfriday/v2"
)

type BlogPost struct {
	Title string
	Body  template.HTML
}

type FileReader func(filename string) ([]byte, error)

func loadBlogPost(title string, fileReader FileReader) (*BlogPost, error) {
	filename := "posts/" + title + ".md"
	body, err := fileReader(filename)
	if err != nil {
		return nil, err
	}
	blogHtml := template.HTML(blackfriday.Run(body))
	return &BlogPost{Title: title, Body: blogHtml}, nil
}

func createPostHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Path[len(("/posts/")):]
		bp, err := loadBlogPost(title, ioutil.ReadFile)
		if err != nil {
			bp = &BlogPost{Title: title, Body: template.HTML("Looks like I never wrote this post!")}
		}
		v.Render(w, r, bp)
	}
}

func createDefaultHandler(v *views.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v.Render(w, r, nil)
	}
}

func main() {
	postView := views.NewView("default", "views/templates/post.html")
	indexView := views.NewView("default", "views/templates/index.html")

	http.HandleFunc("/posts/", createPostHandler(postView))
	http.HandleFunc("/", createDefaultHandler(indexView))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
