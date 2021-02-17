package service

import (
	"HumoAcademy/models"
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

func (s *UserService) CreateUser(user models.Users) (int, error){
	//generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}