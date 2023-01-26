package server

import (
	"net/http"
)

func (s *APIServer) configureRouter() {
	s.Router.HandleFunc("/", s.SendHello()).Methods(http.MethodGet)
}
