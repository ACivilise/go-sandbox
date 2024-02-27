package main

import (
	"log"
	"net/http"

	"github.com/acivilise/go-sandbox/helpers"

	"github.com/acivilise/go-sandbox/handlers/about"
	"github.com/joho/godotenv"

	"github.com/acivilise/go-sandbox/handlers/contact"

	"github.com/acivilise/go-sandbox/handlers/home"

	"github.com/acivilise/go-sandbox/handlers/menu"
	"github.com/acivilise/go-sandbox/handlers/services"
)

func main() {
	// Load .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatalf("Error loading .env file: %v", errEnv)
	}

	// Recursively parse templates
	templates, errTemplates := helpers.ParseTemplates("./templates")
	if errTemplates != nil {
		log.Fatalf("Failed to parse templates: %v", errTemplates)
	}

	// Print the names of all loaded templates to the console
	helpers.PrintTemplateNames(templates)

	http.HandleFunc("/about", about.AboutHandler(templates))
	http.HandleFunc("/contact", contact.ContactHandler(templates))
	http.HandleFunc("/add_item", home.AddItemHandler(templates))
	http.HandleFunc("/remove_item", home.RemoveItemHandler(templates))
	http.HandleFunc("/menu", menu.MenuHandler(templates))
	http.HandleFunc("/services", services.ServicesHandler(templates))
	http.HandleFunc("/", home.HomeHandler(templates))
	http.HandleFunc("/home", home.HomeHandler(templates))

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
