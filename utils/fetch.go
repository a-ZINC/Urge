package utils

import (
	"image"
	_ "image/png"
	"os"
	"urge/log"
	"urge/model"
)

// producer
func ProduceImages(fetchChannel chan<- model.Image, images []model.Image) {
	defer close(fetchChannel)
	for _, img := range images {
		file, err := os.OpenFile(img.Url, os.O_RDONLY, 0766)
		if err != nil {
			log.WarnLogger.Printf("error in opening file: %s", img.Url)
			continue
		}
		decodedImage, fileformat, err := image.Decode(file)
		if err != nil {
			log.WarnLogger.Printf("error in decoding file: %s", img.Url)
			continue
		}
		img.Format = fileformat
		img.Image = decodedImage
		fetchChannel <- img
	}
}
