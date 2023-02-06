package services

import "CustomServerTemplate/internal/repository"

type Service struct {
	UserSrv *UserService
}

func (s *Service) Initialize(db *repository.DB) {
	s.UserSrv = &UserService{
		db.UserRepository,
	}
}
