package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sohail-9098/ms-user-auth/auth"
	"github.com/sohail-9098/ms-user-auth/user"
	"github.com/sohail-9098/ms-user-auth/util"
)

const (
	serverAddress = "localhost:4000"
)

func StartApplication() {
	router := setupRouter()
	log.Println("server started - listening on ", serverAddress)
	if err := http.ListenAndServe(serverAddress, router); err != nil {
		log.Fatalf("could not start the server: %v\n", err)
	}
}

// sets up a gorilla mux router
func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("GET")
	return router
}

// handler func for /login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user user.Credentials

	// decode request body
	if err := decodeBody(r.Body, &user); err != nil {
		log.Printf("error decoding request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error decoding request body: %v", err)
		return
	}

	// validate user fields
	if err := validateUserFields(user); err != nil {
		log.Printf("error validating request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error validating request body: %v", err)
		return
	}

	// hash password
	user.Password = util.HashPassword(user.Password)

	// authenticate against DB
	authorized, err := auth.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		log.Printf("error while authenticating : %v\n", err)
		fmt.Fprintf(w, "error while authenticating: %v", err)
		return
	}
	if !authorized {
		log.Printf("error while authenticating : %v\n", "invalid credentials")
		fmt.Fprintf(w, "error while authenticating: %v", "invalid credentials")
		return
	}

	// create token on successful authentication
	tokenString, err := auth.CreateToken(user.Username, 60)
	if err != nil {
		log.Printf("error while creating token : %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error creating token: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tokenString)
}

// decode request body
func decodeBody(body io.ReadCloser, user *user.Credentials) error {
	if err := json.NewDecoder(body).Decode(&user); err != nil {
		if err.Error() == "EOF" {
			return errors.New("request body should not be empty")
		}
		return err
	}
	return nil
}

// validate user fields
// such as username, password must be present
func validateUserFields(user user.Credentials) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("request missing username or password field")
	}
	return nil
}
