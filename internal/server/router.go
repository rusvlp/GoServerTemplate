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

	// Методы - вебформы
	s.Router.HandleFunc("/register", s.RegisterUser()).Methods(http.MethodGet)
	s.Router.HandleFunc("/register", s.RegisterUser()).Methods(http.MethodPost)

	// Методы - api
	s.Router.HandleFunc("/api/register", s.RegisterUserJSON()).Methods(http.MethodPost)
}
