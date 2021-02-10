package repository

import (
	"HumoAcademy/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MainPagePostgres struct {
	db *sqlx.DB
}

func NewMainPagePostgres(db *sqlx.DB) *MainPagePostgres {
	return &MainPagePostgres{db: db}
}

func (r *MainPagePostgres) GetAll() (models.MainPageContent, error) {
	var Content models.MainPageContent
	queryNews := fmt.Sprintf("SELECT id, title, short_desc, img FROM news")
	err := r.db.Select(&Content.News, queryNews)
	if err != nil {
		return models.MainPageContent{}, err
	}

	queryCourses := fmt.Sprintf("SELECT id, title, course_durance, img FROM courses")
	err = r.db.Select(&Content.Courses, queryCourses)
	if err != nil {
		return models.MainPageContent{}, err
	}

	return Content, nil
}

func (r *MainPagePostgres) AddUserForNews (user models.SubscribedUsers) error {
	query := "INSERT INTO subscribed_users (email) VALUES($1)"
	_, err := r.db.Exec(query, user.Email)
	if err != nil {
		return err
	}
	return nil
}