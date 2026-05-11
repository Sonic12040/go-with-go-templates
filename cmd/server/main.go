package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Application struct {
	tmpl *template.Template
}

func main() {
	// Parse the Templates
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/index.html",
	))

	app := &Application{
		tmpl: tmpl,
	}

	mux := http.NewServeMux()

	// Serve Static Files
	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Routes
	mux.HandleFunc("GET /{$}", app.HandleIndex)
	mux.HandleFunc("POST /ping", app.HandlePing)

	log.Println("Starting Server on port :8080 🚀")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func (app *Application) HandleIndex(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "App Dashboard",
	}

	err := app.tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		log.Printf("Template rendering error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandlePing returns a simple JSON response with the current timestamp
func (app *Application) HandlePing(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":    "success",
		"timestamp": time.Now().Format(time.RFC3339Nano),
		"message":   "Server is running.",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
