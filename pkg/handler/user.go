package handler

import (
	"HumoAcademy/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

const (
	UsersCVDirectory = `images/users_cv/%s_%s`
)

func (h *Handler) getAllSubscribedUsers (c *gin.Context) {
	_ , err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins id param")
		return
	}

	_ , err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins level param")
		return
	}

	emails, err := h.services.User.GetAllSubscribedUsers()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		log.Println(err)
		return
	}
	if emails == nil {
		emails = []string{}
	}
	c.JSON(http.StatusOK, emails)
}

func (h *Handler) SendMail (c *gin.Context) {
	var msg models.MSG
	err := c.BindJSON(&msg)

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad", err.Error())
		return
	}

	message := []byte(
		fmt.Sprintf("Subject : %s", msg.Subject) +
			"\r\n" +
			msg.Message)

	auth := smtp.PlainAuth("", msg.Email, msg.Password, "smtp.mail.ru")
	to, err := h.services.User.GetAllSubscribedUsers()

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		log.Println(err)
		return
	}

	err = smtp.SendMail("smtp.mail.ru:25", auth, msg.Email, to, message)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		log.Println(err)
		return
	}

}

func (h *Handler) createUser (c *gin.Context) {

	cvPath, err := getNewUsersCV(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad", err.Error())
		return
	}

	user, err := getNewUserMainJson(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad", err.Error())
		return
	}
	user.CV = cvPath



	id, err := h.services.User.CreateUser(user)

	if err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, "bad", err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id" : id,
		"status": "ok",
	})
}

func getNewUsersCV(c *gin.Context) (string, error) {
	cv, err := c.FormFile("cv")
	if err != nil {
		log.Println("Error while receiving multipart form. error is", err.Error())
		return "", err
	}

	timeSign := fmt.Sprintf("%d",time.Now().UnixNano())

	cvPath := fmt.Sprintf(UsersCVDirectory, timeSign, cv.Filename)

	file, err := os.Create(cvPath)
	if err != nil {
		log.Println("Error while creating file for cv.", err.Error())
		return "", err
	}
	defer file.Close()

	err = c.SaveUploadedFile(cv, file.Name())
	if err != nil {
		log.Println("Error while saving the cv.", err.Error())
		return "", err
	}
	return cvPath, nil
}

func getNewUserMainJson(c *gin.Context) (models.Users, error) {
	var User models.Users

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error while receiving multipart form. error is", err.Error())
		return models.Users{}, err
	}

	mainJson := form.Value["main_json"]

	err = json.Unmarshal([]byte(mainJson[0]), &User)
	if err != nil {
		log.Println("json unmarshal error:", err.Error())
		return models.Users{}, err
	}

	return User, nil
}

func (h *Handler) getAllCourseUsers (c *gin.Context) {

	_ , err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins id param")
		return
	}

	_ , err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins level param")
		return
	}
	courseIdString := c.Param("course_id")
	courseIdInt, err := strconv.Atoi(courseIdString)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad", err.Error())
		return
	}

	users, err := h.services.User.GetAllCourseUsers(courseIdInt)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}
	if users == nil {
		users = []models.Users{}
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) deleteUserByID(c *gin.Context) {
	_, err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad", "invalid admins id param")
		return
	}

	_, err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad", "invalid admins level param")
		return
	}

	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	NewErrorResponse(c, http.StatusBadRequest, "bad","invalid id param")
	//	return
	//}
}

func (h *Handler) getUserById (c *gin.Context) {

	_ , err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins id param")
		return
	}

	_ , err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins level param")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad","invalid id param")
		return
	}

	course, err := h.services.User.GetUserById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, "bad", err.Error())
		return
	}
	c.JSON(http.StatusOK, course)
}
