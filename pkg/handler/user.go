package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/smtp"
)

func (h *Handler) getAllSubscribedUsers (c *gin.Context) {

	emails, err := h.services.User.GetAllSubscribedUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}
	if emails == nil {
		emails = []string{}
	}
	c.JSON(http.StatusOK, emails)
}


func (h *Handler) SendMail (c *gin.Context) {
	msg := []byte("hello!") //TODO: get from the body of the request

	auth := smtp.PlainAuth("", "namesurname10062001@gmail.ru", "name_surname", "smtp.mail.ru")
	to, err := h.services.User.GetAllSubscribedUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	err = smtp.SendMail("smtp.mail.ru:25", auth, "namesurname10062001@mail.ru", to, msg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		log.Println(err)
	}

}