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

func getNewCourseImg(c *gin.Context) (string, error) {
	img, err := c.FormFile("img")
	if err != nil {
		log.Println("Error while receiving multipart form. error is", err.Error())
		return "", err
	}

	timeSign := fmt.Sprintf("%d",time.Now().UnixNano())

	imgPath := fmt.Sprintf(CoursesImagesDirectory, timeSign, img.Filename)

	file, err := os.Create(imgPath)
	if err != nil {
		fmt.Println("Error while creating file for image.", err.Error())
		return "", err
	}
	err = c.SaveUploadedFile(img, file.Name())
	if err != nil {
		fmt.Println("Error while saving the image.", err.Error())
		return "", err
	}
	return imgPath, nil
}

func getNewCourseMainJson(c *gin.Context) (models.Courses, error) {
	var Course models.Courses

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error while receiving multipart form. error is", err.Error())
		return models.Courses{}, err
	}

	mainJson := form.Value["main_json"]

	err = json.Unmarshal([]byte(mainJson[0]), &Course)
	if err != nil {
		log.Println("json unmarshal error:", err.Error())
		return models.Courses{}, err
	}

	return Course, nil
}

func (h *Handler) createCourse (c *gin.Context) {
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

	imgPath, err := getNewCourseImg(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	course, err := getNewCourseMainJson(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

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

func (h *Handler) getAllCourses (c *gin.Context) {

	courses, err := h.services.Courses.GetAllCourses()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}
	if courses == nil {
		courses = []models.Courses{}
	}
	c.JSON(http.StatusOK, courses)
}

func (h *Handler) deleteCourse(c *gin.Context){
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad","invalid id param")
		return
	}

	err = h.services.Courses.DeleteCourse(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError,"bad", err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message":"course was successfully deleted",
	})
}

func (h *Handler) editCourse (c *gin.Context) {
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "bad","invalid id param")
		return
	}

	imgPath, err := getNewCourseImg(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	course, err := getNewCourseMainJson(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	course.Img = imgPath
	err = h.services.Courses.EditCourse(id, course)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message": "course was successfully updated",
	})
}

func (h *Handler) registerToCourse (c *gin.Context) {

}