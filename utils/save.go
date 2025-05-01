package utils

import (
	"image"
	"image/jpeg"
	"os"
	"strings"
	"time"
	"urge/cmd"
	"urge/log"
)

func Save(img image.Image) {
	directory := strings.Split(cmd.Flags.Input, "/")
	directory = directory[:len(directory)-1]
	directoryString := strings.Join(directory, "/")
	fileLocation := directoryString + "/new.jpeg"

	log.InfoLogger.Printf("----------Saving image------------")
	time.Sleep(2 * time.Second)
	newImageFile, err := os.Create(fileLocation)
	if err != nil {
		log.WarnLogger.Printf("error in creating file: %s", err)
		return
	}
	defer newImageFile.Close()
	err = jpeg.Encode(newImageFile, img, nil)
	if err != nil {
		log.WarnLogger.Printf("error in encoding file: %s", err)
		return
	}
	log.InfoLogger.Printf("image saved at %s", fileLocation)
}
