package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	// Add handler to display "pong" in console
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("pong")
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/priv", privHandler)

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func privHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://go-example-proud-resonance-3596.flycast:8080/ping")
	if err != nil {
		log.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	log.Println("flycast")
	log.Println(time.Now().Format(time.RFC3339))

	w.WriteHeader(http.StatusOK)
}
