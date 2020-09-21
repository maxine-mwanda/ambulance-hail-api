package controllers


import "github.com/gorilla/mux"

func (s *Server) initRoutes() {
	s.Router = mux.NewRouter()

	s.Router.HandleFunc("/", s.Home).Methods("GET")
	s.Router.HandleFunc("/admin", s.Admin).Methods("POST")
	s.Router.HandleFunc("/customer", s.Customer).Methods("POST")
	s.Router.HandleFunc("/driver", s.Driver).Methods("POST")
	s.Router.HandleFunc("/drivers", s.ListDriver).Methods("GET")
	s.Router.HandleFunc("/ambulance", s.Ambulance).Methods("GET")
}