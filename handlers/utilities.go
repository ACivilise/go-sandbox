package handlers

import (
	"html/template"
	"net/http"
)

// Checks if the request is an HTMX request
func IsHXRequest(req *http.Request) bool {
	return req.Header.Get("HX-Request") == "true"
}

func RenderResponse(w http.ResponseWriter, req *http.Request, tmpl *template.Template, fullLayout, partialPath string) {
	var err error
	if IsHXRequest(req) {
		err = tmpl.ExecuteTemplate(w, partialPath, nil)
	} else {
		partial := tmpl.Lookup(partialPath)
		if partial == nil {
			http.Error(w, "Partial template not found", http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteTemplate(w, fullLayout, map[string]interface{}{
			"partial_content": partial,
		})
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
