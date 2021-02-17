package handler

import (
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
		log.Println(err)
		return err
	}
	return nil
}