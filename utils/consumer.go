package utils

import (
	"image"
	"urge/log"
	"urge/model"
)

func ConsumeImages(fetchChannel <-chan model.Image, resizeChannel chan model.Image, filterChannel chan model.Image, saveChannel chan image.Image) {
	defer close(resizeChannel)
	for {
		select {
		case img, ok := <-fetchChannel:
			if !ok {
				log.WarnLogger.Printf("all images has been fetched")
				return
			} else {
				if img.Resize != "" {
					resizeChannel <- img
				} else if img.Filter != "" {
					filterChannel <- img
				} else {
					saveChannel <- img.Image
				}
			}
		}
	}
}

func ConsumeResize(resizeChannel <-chan model.Image, filterChannel chan model.Image, saveChannel chan image.Image) {
	defer close(filterChannel)
	for {
		select {
		case img, ok := <-resizeChannel:
			if !ok {
				log.WarnLogger.Printf("all images has been resized")
				return
			} else {
				Resize(img)
				if img.Filter != "" {
					filterChannel <- img
				} else {
					saveChannel <- img.Image
				}
			}
		}
	}
}

func ConsumeFilter(filterChannel <-chan model.Image, saveChannel chan image.Image) {
	defer close(saveChannel)
	for {
		select {
		case img, ok := <-filterChannel:
			if !ok {
				log.WarnLogger.Printf("all images has been filtered")
				return
			} else {
				filter := Filter(img)
				saveChannel <- filter.GetOutputImage()
			}
		}
	}
}

func ConsumeSave(saveChannel <-chan image.Image) {
	for {
		select {
		case img, ok := <-saveChannel:
			if !ok {
				log.WarnLogger.Printf("all images has been saved")
				return
			} else {
				Save(img)
			}
		}
	}
}
