package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, auth, Accept")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	}

	//c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	//c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Request-With")
	//c.Header("Accept", "*/*")
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("id")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}