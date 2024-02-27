package home

import (
	"html/template"
	"net/http"
)

func RemoveItemHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Define the data to pass to the template
		data := struct {
			Message string
		}{
			Message: "This content was loaded dynamically with HTMX!",
		}

		// Execute the specific template with the data
		err := tmpl.ExecuteTemplate(w, "pages/home/item_add.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	}
}
