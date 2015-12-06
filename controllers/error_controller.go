package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/c0dect/basic-rest-service/models"
)

var (
	ErrBadRequest     = &models.Error{400, "Bad Request", nil}
	ErrInternalServer = &models.Error{500, "Internal Server Error", nil}
)

func WriteError(w http.ResponseWriter, err *models.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(models.Errors{[]*models.Error{err}})
}
