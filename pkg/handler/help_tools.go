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

func deleteImg (c *gin.Context) error {
	imgSrc := c.Query("img_src")
	err := os.Remove(imgSrc)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}