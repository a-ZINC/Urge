package utils

import (
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
	"time"
	"urge/log"
	"urge/model"
)

func Save(img model.Image) {
	_, file := filepath.Split(img.Url)
	file = strings.TrimSuffix(file, filepath.Ext(file))
	fileLocation := filepath.Join(img.OutputUrl, "transformed_" + file + ".jpeg")

	log.InfoLogger.Printf("----------Saving image------------")
	time.Sleep(2 * time.Second)
	newImageFile, err := os.Create(fileLocation)
	if err != nil {
		log.WarnLogger.Printf("error in creating file: %s", err)
		return
	}
	defer newImageFile.Close()
	err = jpeg.Encode(newImageFile, img.Output, nil)
	if err != nil {
		log.WarnLogger.Printf("error in encoding file: %s", err)
		return
	}
	log.InfoLogger.Printf("image saved at %s", fileLocation)
}
