package main

import (
	"fmt"
	"image"
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
	resizeChannel := make(chan model.Image)
	filterChannel := make(chan model.Image)
	saveChannel := make(chan image.Image)
	go utils.ProduceImages(fetchChannel, images)
	go utils.ConsumeImages(fetchChannel, resizeChannel, filterChannel, saveChannel)
	go utils.ConsumeResize(resizeChannel, filterChannel, saveChannel)
	go utils.ConsumeFilter(filterChannel, saveChannel)
	utils.ConsumeSave(saveChannel)
}
