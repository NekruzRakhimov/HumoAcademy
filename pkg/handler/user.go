package handler

import (
	"HumoAcademy/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/smtp"
)

func (h *Handler) getAllSubscribedUsers (c *gin.Context) {
	_ , err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins id param")
		return
	}

	_ , err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins level param")
		return
	}

	emails, err := h.services.User.GetAllSubscribedUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		log.Println(err)
		return
	}
	if emails == nil {
		emails = []string{}
	}
	c.JSON(http.StatusOK, emails)
}


func (h *Handler) SendMail (c *gin.Context) {
	//msg := []byte("hello!") //TODO: get from the body of the request
	//
	//auth := smtp.PlainAuth("", "namesurname10062001@mail.ru", "name_surname", "smtp.mail.ru")
	////to, err := h.services.User.GetAllSubscribedUsers()
	////if err != nil {
	////	newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
	////	log.Println(err)
	////	return
	////}
	//to := []string{"nekruzusa111@gmail.com"}
	//err := smtp.SendMail("smtp.mail.ru:25", auth, "namesurname10062001@mail.ru", to, msg)
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
	//	log.Println(err)
	//	return
	//}
	// Set up authentication information.
	auth := smtp.PlainAuth("", "namesurname10062001@mail.ru", "name_surname", "smtp.mail.ru")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"nekruzusa111@gmail.com"}
	msg := []byte(
		//"To: recipient@example.net\r\n" +
	"Subject: discount Gophers!\r\n" +
	"\r\n" +
	"This is the email body.\r\n")
	err := smtp.SendMail("smtp.mail.ru:25", auth, "namesurname10062001@mail.ru", to, msg)
	if err != nil {
		log.Fatal(err)
	}

}

func (h *Handler) usSignUp (c *gin.Context) {
	var input models.Admin

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad","invalid input body")
		return
	}

	id, err := h.services.Admin.CreateAdmin(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"id": id,
	})
}