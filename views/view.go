package views

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	layoutFiles, err := filepath.Glob("views/layouts/*.html")
	if err != nil {
		panic(err)
	}
	files = append(files, layoutFiles...)
	tmpl := template.Must(template.ParseFiles(files...))
	return &View{
		Template: tmpl,
		Layout:   layout,
	}
}

func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	err := v.Template.ExecuteTemplate(w, v.Layout+".html", data)
	if err != nil {
		fmt.Fprint(w, err.Error(), http.StatusInternalServerError)
	}
}
