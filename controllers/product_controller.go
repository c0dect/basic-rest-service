package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
