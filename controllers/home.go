package controllers


import (
	"github.com/innv8/ambulance-hail-api/middleware"
	"github.com/innv8/ambulance-hail-api/utils"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	utils.Log("welcome home")
	middleware.OkResponse(w, 200, "Ok")
}
