package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello again!")
	})

	wrapped := middlewareHandler(mux)

	log.Println("listening on port : " + port)
	http.ListenAndServe(":"+port, wrapped)
}

func middlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New Request !", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		// apply some logic after call ...
	})
}
