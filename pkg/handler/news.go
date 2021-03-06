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
	NewsImagesDirectory  = `images/news/%s_%s`
)

func getNewsImg(c *gin.Context)  (string, error) {
	img, err := c.FormFile("img")
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad", err.Error())
		return "", err
	}

	timeSign := fmt.Sprintf("%d",time.Now().UnixNano())

	imgPath := fmt.Sprintf(NewsImagesDirectory, timeSign, img.Filename)

	file, err := os.Create(imgPath)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return "", err
	}
	defer file.Close()

	err = c.SaveUploadedFile(img, file.Name())
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return "", err
	}
	return imgPath, nil
}

func getNewsMainJson(c *gin.Context) (models.News, error) {
	var News models.News

	form, err := c.MultipartForm()
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad", err.Error())
		return models.News{}, err
	}

	mainJson := form.Value["main_json"]
	log.Println("main_json", mainJson)

	err = json.Unmarshal([]byte(mainJson[0]), &News)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return models.News{}, err
	}

	return News, nil
}

func (h *Handler) getAllMiniNews (c *gin.Context) {

	news, err := h.services.News.GetAllMiniNews()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError,"bad" ,err.Error())
		return
	}
	if news == nil {
		news = []models.MiniNews{}
	}
	c.JSON(http.StatusOK, news)
}

func (h *Handler) createNews (c *gin.Context) {
	_ , err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins id param")
		return
	}

	_ , err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins level param")
		return
	}

	imgPath, err := getNewsImg(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	news, err := getNewsMainJson(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}
	news.Img = imgPath

	id, err := h.services.News.CreateNews(news)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
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
		NewErrorResponse(c, http.StatusBadRequest, "bad","invalid id param")
		return
	}

	news, err := h.services.News.GetNewsByID(id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, "bad",err.Error())
		return
	}
	c.JSON(http.StatusOK, news)
}

func (h *Handler) editNews (c *gin.Context) {
	_ , err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins id param")
		return
	}

	_ , err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins level param")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad","invalid id param")
		return
	}

	var news models.News
	err = c.BindJSON(&news)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	err = h.services.News.EditNews(id, news)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message": "news was successfully updated",
	})
}

func (h *Handler) changeNewsStatus(c *gin.Context){
	_ , err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins id param")
		return
	}

	_ , err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins level param")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad","invalid id param")
		return
	}

	statusSTR := c.Query("status")
	status := strToBool(statusSTR)

	err = h.services.News.ChangeNewsStatus(id, status)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError,"bad", err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message":"course_status was successfully updated",
	})
}

func (h *Handler) changeNewsImg (c *gin.Context) {
	_ , err := getAdminId(c) //TODO: (adminId) check id
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins id param")
		return
	}

	_ , err = getAdminLevel(c) //TODO: (adminLevel) check for admin level
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "bad","invalid admins level param")
		return
	}
	/*****************************************************************************************************/

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "bad","invalid id param")
		return
	}
	/*****************************************************************************************************/
	imgSrc, err := h.services.News.GetNewsImgSrc(id)
	err = deleteImg(imgSrc)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}
	/*****************************************************************************************************/
	imgPath, err := getNewsImg(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "bad", err.Error())
		return
	}

	err = h.services.News.ChangeNewsImg(id, imgPath)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError,"bad", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
		"message":"news image was successfully updated",
	})
}

