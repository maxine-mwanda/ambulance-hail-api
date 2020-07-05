package controllers

import (
	"encoding/json"
	"github.com/innv8/ambulance-hail-api/entities"
	"github.com/innv8/ambulance-hail-api/middleware"
	"github.com/innv8/ambulance-hail-api/models"
	"net/http"
)

func (s *Server) Customer (w http.ResponseWriter, r *http.Request) {
	var data entities.NewCustomer
	err := json.NewDecoder(r.Body).Decode(&data)
	if err !=nil {
		middleware.ErrorResponse(w, "invalid data")
		return
	}
	err = models.CreateCustomer(data, s.Db)
	if err !=nil {
		middleware.ErrorResponse(w, "unable to save customer")
		return
	}
	middleware.OkResponse(w, 200, data.FirstName)
}
