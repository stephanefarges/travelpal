package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Signup")
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Signin")
}
func RegisterEndpoints(mux *mux.Router) {

	mux.HandleFunc("/user/signup", SignupHandler).Methods("POST")
	mux.HandleFunc("/user/signin", SigninHandler).Methods("POST")
}
