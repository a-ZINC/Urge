package utils

import (
	"fmt"
	"image"
	_ "image/png"
	_ "image/jpeg"
	"net/http"
	"os"
	"strings"
	"urge/log"
	"urge/model"
)

// producer
func ProduceImages(fetchChannel chan<- model.Image, images []model.Image) {
	defer close(fetchChannel)
	var err error
	for _, img := range images {
		if (strings.HasPrefix(img.Url, "http") || strings.HasPrefix(img.Url, "https")) && img.Url != "" {
			err = fetchLink(&img)
		} else {
			err = fetchFile(&img)
		}
		if err != nil {
			log.WarnLogger.Printf("error in fetching image: %s", err)
			continue
		}
		fetchChannel <- img
	}
}

func fetchFile(img *model.Image) error {
	file, err := os.OpenFile(img.Url, os.O_RDONLY, 0766)
	if err != nil {
		return err
	}
	decodedImage, fileformat, err := image.Decode(file)
	if err != nil {
		return err
	}
	img.Format = fileformat
	img.Image = decodedImage
	return nil
}

func fetchLink(img *model.Image) error {
	resp, err := http.Get(img.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("error in fetching image: %s", img.Url)
	}
	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "image") {
		decodedImage, fileformat, err := image.Decode(resp.Body)
		if err != nil {
			return err
		}
		img.Format = fileformat
		img.Image = decodedImage
	}
	return nil
}
