package server

import (
	"net/http"
)

func (s *APIServer) configureRouter() {
	s.Router.HandleFunc("/", s.Index()).Methods(http.MethodGet)
	s.Router.HandleFunc("/testApi", s.TestAPI()).Methods(http.MethodGet)
	s.Router.HandleFunc("/intParam", s.ParamTest()).Methods(http.MethodGet)
}
