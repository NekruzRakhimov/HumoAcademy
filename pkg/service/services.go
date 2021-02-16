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
	EditCourse(id int, course models.Courses) error
	GetCourseById (id int) (newCourse models.Courses, err error)
	GetAllMiniCourses () ([]models.MiniCourses, error)
	ChangeCourseStatus (id int, status bool) error
}

type News interface {
	CreateNews(news models.News) (int, error)
	GetNewsByID (int) (models.News, error)
	GetAllMiniNews () ([]models.MiniNews, error)
	EditNews(id int, news models.News) error
	ChangeNewsStatus (id int, status bool) error
}

type Admin interface {
	CreateAdmin(admin models.Admin) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, int, error)
}

type User interface {
	GetAllSubscribedUsers () ([]string, error)
}

type Service struct {
	MainPage
	Courses
	News
	Admin
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		MainPage: NewMainPageService(repos.MainPage),
		Courses: NewCoursesService(repos.Courses),
		News: NewNewsService(repos.News),
		Admin: NewAdminService(repos.Admin),
		User: NewUserService(repos.User),
	}
}