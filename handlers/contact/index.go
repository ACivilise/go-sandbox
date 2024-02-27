package contact

import (
	"html/template"
	"net/http"

	"github.com/acivilise/go-sandbox/handlers"
)

func ContactHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		const partialPath = "pages/contact/index.html"
		handlers.RenderResponse(w, req, tmpl, "index.html", partialPath)
	}
}
