package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		fmt.Printf("Please provide a valid port...")
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server started in port %s\n", PORT)
	})

	http.ListenAndServe(":"+PORT, nil)
}
