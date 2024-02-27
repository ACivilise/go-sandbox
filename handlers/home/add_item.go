package home

import (
	"html/template"
	"net/http"
)

// AddItem handles the dynamic addition of an item and renders a template with a message
func AddItemHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Define the data to pass to the template
		data := struct {
			Message string
		}{
			Message: "This content was loaded dynamically with HTMX!",
		}

		// Execute the specific template with the data
		err := tmpl.ExecuteTemplate(w, "pages/home/item_remove.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}
