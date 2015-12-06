package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/c0dect/basic-rest-service/models"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

const PRODUCT = "Product"

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	requestProduct := models.Product{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&requestProduct)
	if err != nil {
		responseError := ErrBadRequest
		responseError.Error = err
		WriteError(w, responseError)
		return
	}

	productKey := datastore.NewIncompleteKey(context, PRODUCT, nil)
	productKey, err = datastore.Put(context, productKey, &requestProduct)
	requestProduct.ProductId = productKey.IntID()
	productKey, err = datastore.Put(context, productKey, &requestProduct)
	if err != nil {
		log.Println(err)
		responseError := ErrInternalServer
		responseError.Error = err
		WriteError(w, responseError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(requestProduct); err != nil {
		panic(err)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
