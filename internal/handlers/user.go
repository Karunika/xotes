package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/models"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

func Respond(rw http.ResponseWriter, status int, data interface{}, message string) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(Response{data, message})
}

func RespondPanic(rw http.ResponseWriter, status int, message string) {
	Respond(rw, status, nil, message)
}

func (thisService *HandlerService) GetUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				RespondPanic(rw, http.StatusInternalServerError, "An internal server error occured")
			}
		}()

		user := thisService.DBService.SelectUserByUsername(username)

		Respond(rw, http.StatusOK, user, "successful")
	}
}

func (thisService *HandlerService) SignUpUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				RespondPanic(rw, http.StatusInternalServerError, "An internal server error occured")
			}
		}()
		r.ParseForm()

		var auth models.Auth
		json.NewDecoder(r.Body).Decode(&auth)

		bytePassword := []byte(auth.Password)
		hash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)

		thisService.DBService.InsertUser(auth.Username, auth.Email, string(hash))

		Respond(rw, http.StatusOK, nil, "successful")
	}
}

func (thisService *HandlerService) LoginUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				RespondPanic(rw, http.StatusInternalServerError, "An internal server error occured")
			}
		}()
		r.ParseForm()

		var auth models.Auth
		json.NewDecoder(r.Body).Decode(&auth)

		key := auth.Email
		if auth.Email == "" {
			key = auth.Username
		}
		user := thisService.DBService.SelectUserCredByKey(key)

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))

		if err != nil {
			RespondPanic(rw, http.StatusInternalServerError, "An internal server error occured")
			return
		}

		Respond(rw, http.StatusCreated, user, "User Authentication successful")
	}
}
