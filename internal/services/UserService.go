package services

import (
	"CustomServerTemplate/internal/dto"
	"CustomServerTemplate/internal/repository"
	"net/http"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (service *UserService) CreateJSON(request *http.Request) (*dto.User, error) {
	return nil, nil
}

func (service *UserService) CreateForm(request *http.Request) (*dto.User, error) {
	err := request.ParseForm()
	if err != nil {
		return nil, err
	}

	user := service.UserRepo.GetUserData(request)
	err = service.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) CreateUser(user *dto.User) error {
	err := service.UserRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
