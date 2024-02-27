package about

import (
	"html/template"
	"net/http"

	"github.com/acivilise/go-sandbox/handlers" // Adjust the import path based on your module's name
)

func AboutHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		const partialPath = "pages/about/index.html"
		handlers.RenderResponse(w, req, tmpl, "index.html", partialPath)
	}
}
