package handler

import (
	"HumoAcademy/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getNewsImg(c *gin.Context) string {
	img, err := c.FormFile("img")
	if err != nil {
		log.Println("Error while receiving multipart form. error is", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason" : err.Error(),
		})
		return ""
	}

	timeSign := fmt.Sprintf("%d",time.Now().UnixNano())

	imgPath := fmt.Sprintf("images/resumes/%s_%s", timeSign, img.Filename)

	file, err := os.Create(imgPath)
	if err != nil {
		fmt.Println("Error while creating file for image.", err.Error())
		return ""
	}
	err = c.SaveUploadedFile(img, file.Name())
	if err != nil {
		fmt.Println("Error while saving the image.", err.Error())
		return ""
	}
	return imgPath
}

func getNewsMainJson(c *gin.Context) models.News {
	var News models.News

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error while receiving multipart form. error is", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason" : err.Error(),
		})
		return models.News{}
	}

	mainJson := form.Value["main_json"]
	log.Println("main_json", mainJson)

	err = json.Unmarshal([]byte(mainJson[0]), &News)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "wrong request body",
		})
		log.Println("json unmarshal error:", err.Error())
		return models.News{}
	}

	return News
}

func (h *Handler) createNews (c *gin.Context) {
	_ , err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		return
	}

	_ , err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		return
	}

	imgPath := getNewsImg(c)
	news := getNewsMainJson(c)
	news.Img = imgPath

	id, err := h.services.News.CreateNews(news)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"id" : id,
	})
}

func (h *Handler) getNewsById (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad","invalid id param")
		return
	}

	news, err := h.services.News.GetNewsByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "bad",err.Error())
		return
	}
	c.JSON(http.StatusOK, news)
}
