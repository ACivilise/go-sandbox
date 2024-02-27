package menu

import (
	"html/template"
	"net/http"
)

// Menu renders the menu page
func MenuHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		err := tmpl.ExecuteTemplate(w, "menu.html", nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}
