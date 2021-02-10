package handler

import (
	"HumoAcademy/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) adminSignUp (c *gin.Context) {
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

func (h *Handler) adminSignIn (c *gin.Context) {
	var input models.AdminSignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad", err.Error())
		return
	}

	token, err := h.services.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"status": "ok",
	})
}
