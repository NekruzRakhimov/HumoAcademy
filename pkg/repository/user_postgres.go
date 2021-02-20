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

func (r *UserPostgres) GetAllSubscribedUsers () ([]string, error) {
	var emails []string
	query := fmt.Sprintf("SELECT email FROM subscribed_users")
	err := r.db.Select(&emails, query)
	if err != nil {
		return []string{}, err
	}
	return emails, err
}

func (r *UserPostgres) CreateUser (user models.Users) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO users (first_name, last_name, middle_name, email, about, cv, course_id) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id")

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.MiddleName, user.Email, user.About, user.CV, user.CourseId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetAllCourseUsers (courseId int) ([]models.Users, error) {
	var Users []models.Users
	query := fmt.Sprintf("SELECT id, first_name, last_name, middle_name, email, about, cv, course_id FROM users WHERE course_id=$1 ORDER BY id")
	err := r.db.Select(&Users, query, courseId)
	if err != nil {
		return []models.Users{}, err
	}
	return Users, err
}

//func (r *UserPostgres) GetUserByEmailAndCourseID (email string, courseID int) (models.Users, error) {
//	var user models.Users
//	query := fmt.Sprintf("SELECT id, first_name, last_name, middle_name, email, about, cv, course_id FROM users WHERE email=$1 AND course_id=$2")
//	row := r.db.QueryRow(query, email, courseID)
//	err := row.Scan(&user)
//	if err != nil {
//		return models.Users{}, err
//	}
//	return user, err
//}

func (r *UserPostgres) GetUserById (id int) (models.Users, error) {
	var user models.Users
	mainQuery := fmt.Sprintf("SELECT * FROM users WHERE id=$1")
	err := r.db.Get(&user, mainQuery, id)
	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

func (r *UserPostgres) DeleteUserByID (id int) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id = ($1)")
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}