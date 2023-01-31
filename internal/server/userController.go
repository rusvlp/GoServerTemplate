package server

import (
	"CustomServerTemplate/internal/dto"
	"fmt"
	"net/http"
)

//правильно реализовать конкуррентный доступ к данным (sync.Mutex)

func (s *APIServer) RegisterUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		hw := HandleWrapper{writer: writer}
		user := &dto.User{}
		hw.fromJson(user, request)
		hw.jsonResponse(user)
		fmt.Println(user)
	}
}
