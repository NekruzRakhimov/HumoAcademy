package handler

import (
	"HumoAcademy/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Gets all content for main page
func (h *Handler) getAll(c *gin.Context) {
	content, err := h.services.MainPage.GetAll()

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	if content.News == nil {
		content.News = []models.MiniNews{}
	}

	if content.Courses == nil {
		content.Courses = []models.MiniCourses{}
	}

	c.JSON(http.StatusOK, content)
}

//Adds a new user to the DB for new events
func (h *Handler) addUserForNews(c *gin.Context) {
	var input models.SubscribedUsers
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad", err.Error())
		return
	}

	err := h.services.AddUserForNews(input)
	if err != nil {
		NewErrorResponse(c, http.StatusUnprocessableEntity, "bad","this email already exists")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message": "new user was successfully added to the DB",
	})
}
