package models

type MainPageContent struct {
	News []MiniNews `json:"news"`
	Courses []MiniCourses `json:"courses"`
}


