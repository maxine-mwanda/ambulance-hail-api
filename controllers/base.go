package controllers

import (
	"database/sql"
	"fmt"
	"github.com/innv8/ambulance-hail-api/utils"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Db     *sql.DB
	Router *mux.Router
}

func (s *Server) dbConnect() {
	var err error

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUsername, dbPassword, dbHost, dbName)
	if s.Db, err = sql.Open("mysql", dbURI); err != nil {
		os.Exit(3)
	}

	s.Db.SetMaxIdleConns(64)
	s.Db.SetMaxOpenConns(100)
	s.Db.SetConnMaxLifetime(10 * time.Second)
}

func (s *Server) Init() {
	utils.InitLogger()
	s.dbConnect()
	s.initRoutes()
}

func (s *Server) Run() {
	port := ":" + os.Getenv("PORT")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "content-type", "content-length", "accept-encoding", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})


	if err := http.ListenAndServe(port, handlers.CORS(origins, headers, methods)(s.Router)); err != nil {
		s.Db.Close()
	}
}