package utils

import (
	"urge/model"
)

// producer
func ProduceImages(fetchChannel chan<- model.Image, images []model.Image) {

	
	for _, image := range images {
		fetchChannel <- image
	}
	close(fetchChannel)
}
