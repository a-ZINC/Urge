package utils

import (
	"image"
	"urge/model"
)

func ConsumeImages(fetchChannel <-chan model.Image, resizeChannel chan model.Image, filterChannel chan model.Image, saveChannel chan image.Image) {
	for {
		select {
		case img, ok := <-fetchChannel:
			if !ok {
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
					saveChannel <- img.Image
				}
			}
		}
	}
}

func ConsumeFilter(filterChannel <-chan model.Image, saveChannel chan image.Image) {
	for {
		select {
		case img, ok := <-filterChannel:
			if !ok {
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
				return
			} else {
				Save(img)
			}
		}
	}
}
