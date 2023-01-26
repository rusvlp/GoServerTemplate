package server

import (
	addf "CustomServerTemplate/internal/controller"
	"net/http"
)

func (s *APIServer) configureRouter() {
	s.Router.HandleFunc("/", addf.SendHello()).Methods(http.MethodGet)
}
