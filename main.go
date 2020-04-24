package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Server started")

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

	http.ListenAndServe(":8080", nil)
}
