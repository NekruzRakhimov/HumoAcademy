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
	GetAllMiniCourses() ([]models.MiniCourses, error)
	ChangeCourseStatus (id int, status bool) error
	ChangeCourseImg(id int, img string) error
	GetCourseImgSrc(id int) (string, error)
}

type News interface {
	ChangeNewsImg(id int, img string) error
	CreateNews(news models.News) (int, error)
	GetNewsByID (id int) (models.News, error)
	GetAllMiniNews() ([]models.MiniNews, error)
	EditNews(id int, news models.News) error
	ChangeNewsStatus (id int, status bool) error
	GetNewsImgSrc (id int) (string, error)
	CheckNewsExpireDate(timeAtTheMoment int64) error
}

type Admin interface {
	CreateAdmin(admin models.Admin) (int, error)
	GetAdmin(username, password string) (models.Admin, error)
}

type User interface {
	GetAllSubscribedUsers() ([]string, error)
	CreateUser(user models.Users) (int, error)
	GetAllCourseUsers (courseId int) ([]models.Users, error)
	DeleteUserByID (id int) error
	GetUserById (id int) (models.Users, error)

}

type Repository struct {
	MainPage
	Courses
	News
	Admin
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MainPage : NewMainPagePostgres(db),
		Courses: NewCoursesPostgres(db),
		News: NewNewsPostgres(db),
		Admin: NewAdminPostgres(db),
		User: NewUserPostgres(db),
	}
}
