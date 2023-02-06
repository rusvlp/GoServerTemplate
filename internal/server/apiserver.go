package server

import (
	"CustomServerTemplate/internal/config"
	"CustomServerTemplate/internal/repository"
	"CustomServerTemplate/internal/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	AdminRole          = "admin"
	ctxKeyToken ctxKey = iota
)

type ctxKey int8

type APIServer struct {
	Config   *config.Config
	Router   *mux.Router
	db       *repository.DB
	Services *services.Service
}

func New(config *config.Config) *APIServer {
	return &APIServer{
		Config: config,
		Router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	//testRep()
	err := server.configureDB()
	if err != nil {
		return err
	}
	server.configureRouter()
	server.configureServices()
	return http.ListenAndServe(server.Config.Port, server.Router)
}

func (server *APIServer) configureServices() {
	server.Services = &services.Service{}
	server.Services.Initialize(server.db)
}

func (server *APIServer) configureDB() error {
	db := repository.NewDB(server.Config)
	err := db.Open()
	if err != nil {
		return err
	}
	server.db = db
	server.db.CreateRepositories()
	return nil
}
