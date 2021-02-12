package models

type Courses struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Img string `json:"img"`
	Description string `json:"description"`
	Plans string `json:"plans"`
	CourseDurance int `json:"course_durance" db:"course_durance"`
	Status string `json:"status"`
}

type MiniCourses struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Img string `json:"img"`
	CourseDurance int `json:"course_durance" db:"course_durance"`
}
