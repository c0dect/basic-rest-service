package controllers

import (
	"encoding/json"
	"errors"
	"github.com/c0dect/basic-rest-service/models"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"log"
	"net/http"
	"strconv"
)

const USER = "User"

type TokenAuthentication struct {
	Token string `json:"Token"`
}

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

func crypt(password []byte) ([]byte, error) {
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

	hashedPassword, _ := crypt([]byte(requestUser.Password))
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
	userId, _ := validateUser(context, requestUser)
	if err != nil {
		responseError := ErrUnauthorized
		responseError.Error = err
		WriteError(w, responseError)
		return
	}
	log.Printf(userId)
	jwtToken, err := generateToken(userId)
	tokenObject := TokenAuthentication{}
	tokenObject.Token = jwtToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenObject)
	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}

func validateUser(c context.Context, user models.User) (string, error) {

	log.Println(user.Password)
	query := datastore.NewQuery("User").
		Filter("Username =", user.Username)

	var users []models.User
	keys, _ := query.GetAll(c, &users)
	log.Print(keys)
	log.Print(users)
	log.Print(len(users))
	if len(users) == 1 && bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(user.Password)) == nil {
		log.Println("Found User")
		return strconv.FormatInt(keys[0].IntID(), 10), nil
	}
	return "", errors.New("Unauthorized")
}

func generateToken(userId string) (string, error) {
	jwtConfig := InitJWTConfiguration()
	return jwtConfig.GenerateToken(userId)
}
