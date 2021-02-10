package service

import (
	"HumoAcademy/models"
	"HumoAcademy/pkg/repository"
)


type MainPage interface {
	GetAll () (models.MainPageContent, error)
	AddUserForNews (news models.SubscribedUsers) error
}

type Courses interface {
	GetCourseById (id int) (newCourse models.Courses, err error)
	CreateCourse(courses models.Courses) (int, error)
}

type News interface {
	GetNewsByID (int) (models.News, error)
}

type Admin interface {
	CreateAdmin(admin models.Admin) (int, error)
	GenerateToken(username, password string) (string, error)
	//ParseToken(token string) (int, error)
}

type Service struct {
	MainPage
	Courses
	News
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		MainPage: NewMainPageService(repos.MainPage),
		Courses: NewCoursesService(repos.Courses),
		News: NewNewsService(repos.News),
		Admin: NewAdminService(repos.Admin),
	}
}