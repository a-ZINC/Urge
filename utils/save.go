package utils

import (
	"time"
	"urge/log"
	"urge/model"
)

func Save(img model.Image) {
	log.InfoLogger.Printf("----------Saving image------------")
	time.Sleep(2 * time.Second)
}
