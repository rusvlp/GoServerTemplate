package server

import (
	"CustomServerTemplate/internal/dto"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

const userConfigPath string = "config/userConfig"

//правильно реализовать конкурентный доступ к данным (sync.Mutex)

func (s *APIServer) RegisterUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		hw := HandleWrapper{writer: writer}
		if request.Method == http.MethodGet {
			hw.htmlResponse("./web/userSignUp.html")
		}
		if request.Method == http.MethodPost {
			user, err := s.Services.UserSrv.CreateForm(request)
			if err != nil {
				logrus.Error(err)
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			writer.WriteHeader(200)
			fmt.Fprintf(writer, "Пользователь %s успешно создан!", user.Username)
		}

	}
}

func (s *APIServer) RegisterUserJSON() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		hw := HandleWrapper{writer: writer}
		user := dto.User{}
		hw.fromJson(&user, request)
		err := s.Services.UserSrv.CreateUser(&user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		result := map[string]string{
			"Created User: ": user.Username,
		}
		hw.jsonResponse(result)
	}
}
