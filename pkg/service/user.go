package service

import (
	"HumoAcademy/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}

}

func (s *UserService) GetAllSubscribedUsers () ([]string, error) {
	return s.repo.GetAllSubscribedUsers()
}

