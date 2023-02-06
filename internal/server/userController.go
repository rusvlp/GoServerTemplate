package server

import (
	"fmt"
	"net/http"
)

const userConfigPath string = "config/userConfig"

//правильно реализовать конкуррентный доступ к данным (sync.Mutex)

func (s *APIServer) RegisterUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Варианты для JSON и WebForm

		if request.Method == http.MethodGet {
			SendHTML(writer, "./web/userSignUp.html")
		}
		if request.Method == http.MethodPost {
			err := s.Services.UserSrv.CreateForm(request)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			writer.WriteHeader(200)
			fmt.Fprintf(writer, "Пользователь успешно создан!")
		}

	}
}
