package home

import (
	"html/template"
	"net/http"

	"github.com/acivilise/go-sandbox/handlers" // Adjust the import path based on your module's name
)

func HomeHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		const partialPath = "pages/home/index.html"
		handlers.RenderResponse(w, req, tmpl, "index.html", partialPath)
	}
}
