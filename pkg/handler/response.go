package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, status, message string) {
	log.Printf(message, "")
	c.AbortWithStatusJSON(statusCode, statusResponse{status, message})
}
