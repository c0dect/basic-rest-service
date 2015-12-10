package controllers

import (
	"encoding/json"
	"github.com/c0dect/basic-rest-service/models"
	"log"
	"net/http"
)

var (
	ErrBadRequest     = &models.Error{400, "Bad Request", nil}
	ErrInternalServer = &models.Error{500, "Internal Server Error", nil}
	ErrUnauthorized   = &models.Error{401, "Unauthorized", nil}
)

func WriteError(w http.ResponseWriter, err *models.Error) {
	w.Header().Set("Content-Type", "application/json")
	log.Println(err.Error)
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(models.Errors{[]*models.Error{err}})
}
