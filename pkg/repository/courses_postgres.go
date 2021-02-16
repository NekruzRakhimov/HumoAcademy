package repository

import (
	"HumoAcademy/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CoursesPostgres struct {
	db *sqlx.DB
}

func NewCoursesPostgres(db *sqlx.DB) *CoursesPostgres {
	return &CoursesPostgres{db: db}
}

func (r *CoursesPostgres) CreateCourse(course models.Courses) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO courses (title, img, description, plans, course_durance, status) VALUES($1, $2, $3, $4, $5, $6) RETURNING id")
	row := r.db.QueryRow(query, course.Title, course.Img, course.Description, course.Plans, course.CourseDurance, course.Status)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CoursesPostgres) EditCourse(id int, course models.Courses) error {
	query := fmt.Sprintf("UPDATE courses SET title=$1, img=$2, description=$3, plans=$4, course_durance=$5, status=$6 WHERE id=$7")

	_, err := r.db.Exec(query, course.Title, course.Img, course.Description, course.Plans, course.CourseDurance, course.Status, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *CoursesPostgres) GetCourseById (id int) (models.Courses, error) {
	var course models.Courses
	mainQuery := fmt.Sprintf("SELECT * FROM courses WHERE id=$1")
	err := r.db.Get(&course, mainQuery, id)
	if err != nil {
		return models.Courses{}, err
	}

	return course, nil
}

func (r *CoursesPostgres) GetAllMiniCourses() ([]models.MiniCourses, error) {
	var courses []models.MiniCourses
	query := fmt.Sprintf("SELECT id, title, img, course_durance, status FROM courses")
	err := r.db.Select(&courses, query)
	if err != nil {
		return []models.MiniCourses{}, err
	}
	return courses, err
}

func (r *CoursesPostgres) ChangeCourseStatus (id int, status bool) error {
	query := fmt.Sprintf("UPDATE courses SET status = $1 where id = $2")
		_, err := r.db.Exec(query, status, id)
		if err != nil {
			return err
		}
		return nil
}