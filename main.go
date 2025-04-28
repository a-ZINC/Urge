package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"urge/cmd"
	"urge/log"
	"urge/model"
	"urge/utils"
)

func main() {
	cmd.Execute()

	images, err := utils.InputParser()
	fmt.Printf("images: %v \n", images);
	if err != nil {
		log.ErrorLogger.Printf("error: %s", err)
	}
	fetchChannel := make(chan model.Image)
	for _, img := range images {
		go func() {
			file, err := os.Open(img.Url)
			if err != nil {
				log.ErrorLogger.Printf("failed to fetch image: %s", img.Url)
				return
			}
			defer file.Close()
			_image, format, err := image.Decode(file)
			if err != nil {
				log.ErrorLogger.Printf("error decoding image")
			}
			img.Image = _image
			img.Format = format
			fetchChannel <- img
		}()
	}

	for val := range fetchChannel {
		log.InfoLogger.Printf("val: %v", val)
		file, err := os.OpenFile("modified_something.png", os.O_CREATE | os.O_APPEND, 0755)
		if err != nil {
			log.ErrorLogger.Printf("error in creating new file")
		}
		err = png.Encode(file, val.Image)
		if err != nil {
			log.ErrorLogger.Printf("error in creating new image")
		}
	}

}
