package controllers

import (
	"encoding/json"
	"github.com/innv8/ambulance-hail-api/entities"
	"github.com/innv8/ambulance-hail-api/middleware"
	"github.com/innv8/ambulance-hail-api/models"
	"net/http"
)

func (s *Server) Ambulance (w http.ResponseWriter, r *http.Request) {
	var data entities.NewAmbulance
	err := json.NewDecoder(r.Body).Decode(&data)
	if err !=nil {
		middleware.ErrorResponse(w, "invalid data")
		return
	}
	err = models.CreateAmbulance(data, s.Db)
	if err !=nil {
		middleware.ErrorResponse(w, "unable to save ambulance")
		return
	}
	middleware.OkResponse(w, 200, data.CarModel)
}
