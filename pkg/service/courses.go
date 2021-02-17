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

func (s *CoursesService) EditCourse(id int, course models.Courses) error {
	return s.repo.EditCourse(id, course)
}

func (s *CoursesService) GetCourseById (id int) (Course models.Courses, err error) {
	return s.repo.GetCourseById(id)
}

func (s *CoursesService) GetAllMiniCourses () ([]models.MiniCourses, error) {
	return s.repo.GetAllMiniCourses()
}

func (s *CoursesService) ChangeCourseStatus (id int, status bool) error {
	return s.repo.ChangeCourseStatus(id, status)
}

func (s *CoursesService) ChangeCourseImg(id int, img string) error {
	return s.repo.ChangeCourseImg(id, img)
}

func (s *CoursesService) GetCourseImgSrc(id int) (string, error) {
	return s.repo.GetCourseImgSrc(id)
}