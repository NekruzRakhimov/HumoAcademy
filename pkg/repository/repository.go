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
	EditCourse(id int, course models.Courses) error
	GetCourseById (int) (models.Courses, error)
	GetAllCourses() ([]models.Courses, error)
	DeleteCourse (id int) error

}

type News interface {
	CreateNews(news models.News) (int, error)
	GetNewsByID (id int) (models.News, error)
	GetAllNews() ([]models.News, error)
	DeleteNews (id int) error
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
