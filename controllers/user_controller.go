package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/c0dect/basic-rest-service/models"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

const USER = "User"

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

func Crypt(password []byte) ([]byte, error) {
	defer clear(password)
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func Register(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	requestUser := models.User{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&requestUser)
	if err != nil {
		responseError := ErrBadRequest
		responseError.Error = err
		WriteError(w, responseError)
		return
	}

	hashedPassword, _ := Crypt([]byte(requestUser.Password))
	requestUser.Password = string(hashedPassword)

	userKey := datastore.NewIncompleteKey(context, USER, nil)
	userKey, err = datastore.Put(context, userKey, &requestUser)
	if err != nil {
		responseError := ErrInternalServer
		responseError.Error = err
		WriteError(w, responseError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(requestUser); err != nil {
		panic(err)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Logout(w http.ResponseWriter, r *http.Request) {

}
