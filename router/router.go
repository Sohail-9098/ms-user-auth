package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sohail-9098/ms-user-auth/auth"
	"github.com/sohail-9098/ms-user-auth/user"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error decoding request body: %v", err)
		return
	}

	if auth.AuthenticateUser(u.Username, u.Password) {
		tokenString, err := auth.CreateToken(u.Username, 60)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error creating token: %v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, "invalid credentials")
}

func StartApplication() {
	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("POST")

	log.Println("server starting.. listening on port 4000")
	err := http.ListenAndServe("localhost:4000", router)
	if err != nil {
		fmt.Println("could not start the server", err)
	}
}
