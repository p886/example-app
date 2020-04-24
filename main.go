package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Missing env variable 'PORT'")
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/date", func(w http.ResponseWriter, r *http.Request) {
		currentDate := time.Now()
		response := fmt.Sprintf(
			"The current time is %s\nReload this page to get an update\n", currentDate.String(),
		)
		w.Write([]byte(response))
	})

	log.Printf("Server started, listening on port: %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
