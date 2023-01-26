package controller

import (
	"net/http"
)

func SendHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte("Hello world!"))

		return
	}
}

//func (server *APIServer) Index() http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		// Тут нужно вернуть главную страницу
//	}
//}
