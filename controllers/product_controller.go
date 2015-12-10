package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/c0dect/basic-rest-service/dal"
	"github.com/c0dect/basic-rest-service/models"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

func Index(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Fprint(w, "Welcome!\n")
}

func CreateProduct(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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

	productDAL := dal.NewProductDAL(context)
	createdProduct, err := productDAL.AddProduct(requestProduct)

	if err != nil {
		log.Println(err)
		responseError := ErrInternalServer
		responseError.Error = err
		WriteError(w, responseError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdProduct); err != nil {
		panic(err)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context := appengine.NewContext(r)

	productDAL := dal.NewProductDAL(context)
	products, err := productDAL.GetProducts()
	if err != nil {
		log.Println(err)
		responseError := ErrInternalServer
		responseError.Error = err
		WriteError(w, responseError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(products); err != nil {
		panic(err)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context := appengine.NewContext(r)
	productId := mux.Vars(r)["productId"]

	productDAL := dal.NewProductDAL(context)
	product, err := productDAL.GetProduct(productId)
	if err != nil {
		log.Println(err)
		responseError := ErrInternalServer
		responseError.Error = err
		WriteError(w, responseError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		panic(err)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	context := appengine.NewContext(r)
	productId := mux.Vars(r)["productId"]

	productDAL := dal.NewProductDAL(context)
	product, err := productDAL.DeleteProduct(productId)
	if err != nil {
		log.Println(err)
		responseError := ErrInternalServer
		responseError.Error = err
		WriteError(w, responseError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		panic(err)
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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

	productId := mux.Vars(r)["productId"]

	productDAL := dal.NewProductDAL(context)
	product, err := productDAL.UpdateProduct(productId, requestProduct)
	if err != nil {
		//log.Println(err)
		responseError := ErrInternalServer
		responseError.Error = err
		WriteError(w, responseError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		panic(err)
	}
}
