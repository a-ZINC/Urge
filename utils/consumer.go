package utils

import (
	"urge/log"
	"urge/model"
)

func ConsumeImages(fetchChannel <-chan model.Image, resizeChannel chan model.Image, filterChannel chan model.Image, saveChannel chan model.Image) {
	defer close(resizeChannel)
	for {
		select {
		case img, ok := <-fetchChannel:
			if !ok {
				log.WarnLogger.Printf("all images has been fetched")
			} else {
				if img.Resize != "" {
					resizeChannel <- img
				} else if img.Filter != "" {
					filterChannel <- img
				} else {
					saveChannel <- img
				}
			}
		}
	}
}

func ConsumeResize(resizeChannel <-chan model.Image, filterChannel chan model.Image, saveChannel chan model.Image) {
	defer close(filterChannel)
	for {
		select {
		case img, ok := <-resizeChannel:
			if !ok {
				log.WarnLogger.Printf("all images has been resized")
			} else {
				Resize(img)
				if img.Filter != "" {
					filterChannel <- img
				} else {
					saveChannel <- img
				}
			}
		}
	}
}

func ConsumeFilter(filterChannel <-chan model.Image, saveChannel chan model.Image) {
	defer close(saveChannel)
	for {
		select {
		case img, ok := <-filterChannel:
			if !ok {
				close(saveChannel)
				log.WarnLogger.Printf("all images has been filtered")
			} else {
				Filter(img)
				saveChannel <- img
			}
		}
	}
}

func ConsumeSave(saveChannel <-chan model.Image) {
	for {
		select {
		case img, ok := <-saveChannel:
			if !ok {
				log.WarnLogger.Printf("all images has been saved")
			} else {
				Save(img)
			}
			return
		}
	}
}
