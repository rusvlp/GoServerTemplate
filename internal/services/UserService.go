package services

import (
	"CustomServerTemplate/internal/repository"
	"net/http"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (service *UserService) CreateJSON(request *http.Request) error {
	return nil
}

func (service *UserService) CreateForm(request *http.Request) error {
	err := request.ParseForm()
	if err != nil {
		return err
	}

	user := service.UserRepo.GetUserData(request)
	err = service.UserRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
