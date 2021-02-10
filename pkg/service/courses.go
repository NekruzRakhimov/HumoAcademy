package service

import (
	"HumoAcademy/models"
	"HumoAcademy/pkg/repository"
)

const (
	dividingSign = "@"
)

type CoursesService struct {
	repo repository.Courses
}

func NewCoursesService(repo repository.Courses) *CoursesService {
	return &CoursesService{repo: repo}
}

func (s *CoursesService) CreateCourse(course models.Courses) (int, error) {
	return s.repo.CreateCourse(course)
}

func (s *CoursesService) GetCourseById (id int) (Course models.Courses, err error) {
	return s.repo.GetCourseById(id)
}