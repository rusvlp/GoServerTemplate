package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

const userConfigPath string = "config/userConfig"

//правильно реализовать конкуррентный доступ к данным (sync.Mutex)

func (s *APIServer) RegisterUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//hw := HandleWrapper{writer: writer}
		//user := &dto.User{}
		//hw.fromJson(user, request)

		request.ParseForm()

		user := s.db.UserRepository.GetUserData(request)
		err := s.db.UserRepository.Create(user)

		if err != nil {
			SendHTML(writer, "./web/userSignUp.html")
			logrus.Error(err)
		}

		writer.WriteHeader(200)
		fmt.Fprintf(writer, "Пользователь успешно создан! Username: %s, Passowrd: %s", user.Username, user.Password)
		fmt.Println(user)
	}
}
