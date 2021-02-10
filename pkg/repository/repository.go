package repository

import (
	"HumoAcademy/models"
	"github.com/jmoiron/sqlx"
)

type MainPage interface {
	GetAll()(models.MainPageContent, error)
	AddUserForNews (users models.SubscribedUsers) error
}

type Courses interface {
	CreateCourse(course models.Courses) (int, error)
	GetCourseById (int) (models.Courses, error)
}

type News interface {
	GetNewsByID (id int) (models.News, error)
}

type Admin interface {
	CreateAdmin(admin models.Admin) (int, error)
	GetAdmin(username, password string) (models.Admin, error)
}

type Repository struct {
	MainPage
	Courses
	News
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MainPage : NewMainPagePostgres(db),
		Courses: NewCoursesPostgres(db),
		News: NewNewsPostgres(db),
		Admin: NewAdminPostgres(db),
	}
}
