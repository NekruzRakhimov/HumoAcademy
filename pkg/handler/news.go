package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getNewsById (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	news, err := h.services.News.GetNewsByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, news)
}
