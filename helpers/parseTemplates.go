package helpers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ParseTemplates(rootDir string) (*template.Template, error) {
	tmpl := template.New("")
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			// Read the content of the HTML file
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			// Define a template name based on its relative path from rootDir
			relPath, err := filepath.Rel(rootDir, path)
			if err != nil {
				return err
			}
			name := strings.Replace(relPath, "\\", "/", -1) // Ensure the name uses forward slashes
			// Parse the template content and associate it with the name
			_, err = tmpl.New(name).Parse(string(content))
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func PrintTemplateNames(tmpl *template.Template) {
	templates := tmpl.Templates()
	for _, t := range templates {
		// Each template has a Name() method that returns the template's name
		// In this context, it's the relative path used when parsing the templates
		fmt.Println("Loaded template:", t.Name())
	}
}
