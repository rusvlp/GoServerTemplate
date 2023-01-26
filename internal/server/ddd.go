package server

import "net/http"

func (server *APIServer) Index() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Тут нужно вернуть главную страницу
	}
}
