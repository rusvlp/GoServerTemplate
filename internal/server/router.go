package server

import (
	"net/http"
)

func (s *APIServer) configureRouter() {
	s.Router.HandleFunc("/", s.SendHello()).Methods(http.MethodGet)
	s.Router.HandleFunc("/testapi", s.TestAPI()).Methods(http.MethodGet)
}
