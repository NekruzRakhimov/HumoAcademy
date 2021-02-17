package repository

import (
	"HumoAcademy/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAllSubscribedUsers() ([]string, error) {
	var emails []string
	query := fmt.Sprintf("SELECT email FROM subscribed_users")
	err := r.db.Select(&emails, query)
	if err != nil {
		return []string{}, err
	}
	return emails, err
}

func (r *UserPostgres) CreateUser(user models.Users) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO applicants_users (first_name, last_name, middle_name, email, about, cv, course_id) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id")

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.MiddleName, user.Email, user.About, user.CV, user.CourseId)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}


