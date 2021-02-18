package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func strToBool(s string) bool {
	if s == "true" {
		return true
	}

	return false
}

func deleteImg (imgSrc string) error {
	err := os.Remove(imgSrc)
	if err != nil {
		log.Println("Error: while deleting image. Error is ", err)
		return err
	}
	return nil
}

func (h *Handler) getImg(c *gin.Context) {
	imgPath := c.Param("img")
	//f1 := "images"
	f2 := c.Param("f2")

	c.File("./images/" + f2 + "/" + imgPath)
}