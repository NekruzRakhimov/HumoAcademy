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
	CreateCourse(courses models.Courses) (int, error)
	GetCourseById (id int) (newCourse models.Courses, err error)
	GetAllCourses () ([]models.Courses, error)
}

type News interface {
	CreateNews(news models.News) (int, error)
	GetNewsByID (int) (models.News, error)
	GetAllNews () ([]models.News, error)
}

type Admin interface {
	CreateAdmin(admin models.Admin) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, int, error)
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