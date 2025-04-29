package utils

import (
	"time"
	"urge/log"
	"urge/model"
)

func Resize(img model.Image) {
	log.InfoLogger.Printf("----------Resizing image------------")
	time.Sleep(2 * time.Second)
}
