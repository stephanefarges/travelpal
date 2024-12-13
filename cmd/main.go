package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"travelpal/user"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := mux.NewRouter()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello again!")
	})

	user.RegisterEndpoints(mux)

	log.Println("listening on port : " + port)
	http.ListenAndServe(":"+port, mux)
}
