package server

import (
	"CustomServerTemplate/internal/config"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	AdminRole          = "admin"
	ctxKeyToken ctxKey = iota
)

type ctxKey int8

type APIServer struct {
	Config *config.Config
	Router *mux.Router
}

func New(config *config.Config) *APIServer {
	return &APIServer{
		Config: config,
		Router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	server.configureRouter()
	return http.ListenAndServe(server.Config.Port, server.Router)
}
