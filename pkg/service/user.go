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

func (s *UserService) GetAllCourseUsers (courseId int) (models.CourseUsersList, error) {
	return s.repo.GetAllCourseUsers(courseId)
}

func (s *UserService)  DeleteUserByID (id int) error {
	return s.repo.DeleteUserByID(id)
}

func (s *UserService) GetUserById (id int) (User models.Users, err error) {
	return s.repo.GetUserById(id)
}
