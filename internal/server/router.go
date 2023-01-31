package server

import (
	"net/http"
)

func (s *APIServer) configureRouter() {
	// Тестовые методы
	s.Router.HandleFunc("/", s.Index()).Methods(http.MethodGet)
	s.Router.HandleFunc("/testApi", s.TestAPI()).Methods(http.MethodGet)
	s.Router.HandleFunc("/intParam/", s.ParamTest()).Methods(http.MethodGet)
	s.Router.HandleFunc("/pathParam/{id}", s.PathParamTest()).Methods(http.MethodGet)

	// Методы для юзера
	s.Router.HandleFunc("/register", s.RegisterUser()).Methods(http.MethodPost)

}
