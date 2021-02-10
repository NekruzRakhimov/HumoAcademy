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

const (
	CoursesImagesDirectory  = `images/courses/%s_%s`
)

func getNewCourseImg(c *gin.Context) string {
	img, err := c.FormFile("img")
	if err != nil {
		log.Println("Error while receiving multipart form. error is", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
			"status" : "bad",
		})
		return ""
	}

	timeSign := fmt.Sprintf("%d",time.Now().UnixNano())

	imgPath := fmt.Sprintf(CoursesImagesDirectory, timeSign, img.Filename)

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

func geNewCourseMainJson(c *gin.Context) models.Courses {
	var Course models.Courses

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error while receiving multipart form. error is", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
			"status" : "bad",
		})
		return models.Courses{}
	}

	mainJson := form.Value["main_json"]
	log.Println("main_json", mainJson)

	err = json.Unmarshal([]byte(mainJson[0]), &Course)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
			"status" : "bad",
		})
		log.Println("json unmarshal error:", err.Error())
		return models.Courses{}
	}

	return Course
}

func (h *Handler) createCourse (c *gin.Context) {
	_ , err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		return
	}

	_ , err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		return
	}

	imgPath := getNewCourseImg(c)
	course := geNewCourseMainJson(c)
	course.Img = imgPath

	id, err := h.services.Courses.CreateCourse(course)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id" : id,
		"status": "ok",
	})
}

func (h *Handler) getCourseById (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad","invalid id param")
		return
	}

	course, err := h.services.Courses.GetCourseById(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "bad", err.Error())
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *Handler) registerToCourse (c *gin.Context) {

}