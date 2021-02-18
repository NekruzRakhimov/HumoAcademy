package service

import (
	"HumoAcademy/models"
	"HumoAcademy/pkg/repository"
	"strings"
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
	FLM := strings.Split(user.FirstName, " ") //FLM = ФИО
	user.FirstName = FLM[0]
	user.LastName = FLM[1]
	user.MiddleName = FLM[2]
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers () ([]models.Users, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService)  DeleteUserByID (id int) error {
	return s.repo.DeleteUserByID(id)
}

func (s *UserService) GetUserById (id int) (User models.Users, err error) {
	return s.repo.GetUserById(id)
}
