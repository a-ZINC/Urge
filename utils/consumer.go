package utils

import (
	"urge/log"
	"urge/model"
)

func ConsumeImages(fetchChannel <-chan model.Image, rotateChannel chan model.Image, resizeChannel chan model.Image, filterChannel chan model.Image, saveChannel chan model.Image) {
	for {
		select {
		case img, ok := <-fetchChannel:
			if !ok {
				return
			} else {
				if img.Rotate != 0 {
					rotateChannel <- img
				} else if img.Resize != "" {
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

func ConsumeRotate(rotateChannel <-chan model.Image, resizeChannel chan model.Image, filterChannel chan model.Image, saveChannel chan model.Image) {
	for img := range rotateChannel {
		rotateImage, err := Rotate(img)
		if err != nil {
			log.WarnLogger.Printf("error in rotating image: %s", err)
			continue
		}
		if img.Resize != "" {
			resizeChannel <- rotateImage
		} else if img.Filter != "" {
			filterChannel <- rotateImage
		} else {
			saveChannel <- rotateImage
		}
	}
}

func ConsumeResize(resizeChannel <-chan model.Image, filterChannel chan model.Image, saveChannel chan model.Image) {
	for {
		select {
		case img, ok := <-resizeChannel:
			if !ok {
				return
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
	for {
		select {
		case img, ok := <-filterChannel:
			if !ok {
				return
			} else {
				filter, err := Filter(img)
				if err != nil {
					log.WarnLogger.Printf("error in filtering image: %s", err)
					continue
				}
				saveChannel <- filter.GetImage()
			}
		}
	}
}

func ConsumeSave(saveChannel <-chan model.Image) {
	for {
		select {
		case img, ok := <-saveChannel:
			if !ok {
				return
			} else {
				Save(img)
			}
		}
	}
}
