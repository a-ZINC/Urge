package utils

import (
	"fmt"
	"urge/controller/rotate"
	"urge/log"
	"urge/model"
)

func Rotate(img model.Image) (model.Image, error) {
	log.InfoLogger.Printf("----------Rotating image------------")
	rotate := rotate.New(img)
	switch {
	case img.Rotate == 90 || img.Rotate == -270:
		rotate.Rotate90()
	case img.Rotate == 180 || img.Rotate == -180:
		rotate.Rotate180()
	case img.Rotate == 270 || img.Rotate == -90:
		rotate.Rotate270()
	default:
		return rotate.GetImage(), fmt.Errorf("we only support 90, 180 and 270 degree rotations")
	}
	return rotate.GetImage(), nil
}
