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

func getImg(c *gin.Context) string {
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

func getMainJson(c *gin.Context) models.Courses {
	var Course models.Courses

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error while receiving multipart form. error is", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"reason" : err.Error(),
		})
		return models.Courses{}
	}

	mainJson := form.Value["main_json"]
	log.Println("main_json", mainJson)

	err = json.Unmarshal([]byte(mainJson[0]), &Course)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"reason": "wrong request body",
		})
		log.Println("json unmarshal error:", err.Error())
		return models.Courses{}
	}

	return Course
}

func (h *Handler) newCourse (c *gin.Context) {
	//TODO:checking for being admin

	imgPath := getImg(c)
	course := getMainJson(c)
	course.Img = imgPath

	//if err := c.BindJSON(&course); err != nil {
	//	newErrorResponse(c, http.StatusBadRequest, err.Error())
	//	return
	//}

	id, err := h.services.Courses.CreateCourse(course)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id" : id,
	})
}

func (h *Handler) getCourseById (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	course, err := h.services.Courses.GetCourseById(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *Handler) registerToCourse (c *gin.Context) {

}