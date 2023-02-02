package server

import (
	"CustomServerTemplate/internal/dto"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

//правильно реализовать конкуррентный доступ к данным (sync.Mutex)

func (s *APIServer) RegisterUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		hw := HandleWrapper{writer: writer}
		user := &dto.User{}
		hw.fromJson(user, request)
		user, err := s.db.UserRepository.Create(user)
		//hw.jsonResponse(user)
		if err != nil {
			http.Error(writer, "error creating user", http.StatusInternalServerError)
			logrus.Error(err)
		}
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "Пользователь успешно создан! Username: %s, Passowrd: %s", user.Username, user.Password)
		fmt.Println(user)
	}
}
