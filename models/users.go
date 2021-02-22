package models

type Users struct {
	Id int `json:"id"`
	FullName string `json:"full_name" db:"full_name"`
	Email string `json:"email"`
	About string `json:"about"`
	CV string `json:"cv"`
	CourseId int `json:"course_id,omitempty" db:"course_id"`
	//Password string `json:"password" db:"password_hash"`
}

type SubscribedUsers struct {
	Id int `json:"id"`
	Email string `json:"email"`
}

type MSG struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Subject string `json:"subject"`
	Message string `json:"message"`
} 
