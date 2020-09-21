package controllers

import (
	"encoding/json"
	"github.com/innv8/ambulance-hail-api/entities"
	"github.com/innv8/ambulance-hail-api/middleware"
	"github.com/innv8/ambulance-hail-api/models"
	"net/http"
)

func(s *Server) ListDriver (w http.ResponseWriter, r *http.Request) {
	

}

func (s *Server) Driver (w http.ResponseWriter, r *http.Request) {
	var data entities.NewDriver
	err := json.NewDecoder(r.Body).Decode(&data)
	if err !=nil {
		middleware.ErrorResponse(w, "invalid data")
		return
	}
	err = models.CreateDriver(data, s.Db)
	if err !=nil {
		middleware.ErrorResponse(w, "unable to save driver")
		return
	}
	middleware.OkResponse(w, 200, data.FirstName)
}

