package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	adminAuthorizationHeader = "Authorization"
	adminIdCtx = "adminId"
	adminLevelCtx = "adminLevel"
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

func (h *Handler) adminIdentity(c *gin.Context) {
	header := c.GetHeader(adminAuthorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "bad","empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "bad","invalid auth header")
		return
	}

	adminId, adminLevel, err := h.services.Admin.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "bad", err.Error())
	}
	c.Set(adminIdCtx, adminId)
	c.Set(adminLevelCtx, adminLevel)
}

func getAdminId(c *gin.Context) (int, error) {
	adminId, ok := c.Get(adminIdCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "bad","user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := adminId.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "bad","user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}

func getAdminLevel(c *gin.Context) (int, error) {
	adminLevel, ok := c.Get(adminLevelCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "bad","user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := adminLevel.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "bad","user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}