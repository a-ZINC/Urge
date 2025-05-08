package utils

import (
	"fmt"
	"strings"
	"urge/controller/flip"
	"urge/log"
	"urge/model"
)

func Flip(img model.Image) (*model.Image, error) {
	log.InfoLogger.Printf("----------Flipping image------------")
	flipObject := flip.New(img)
	if img.Flip == "h" || img.Flip == "H" || strings.ToLower(img.Flip) == "horizontal" {
		flipObject.FlipH()
	} else if img.Flip == "v" || img.Flip == "V" || strings.ToLower(img.Flip) == "vertical" {
		flipObject.FlipV()
	}
	switch {
	case img.Flip == "h" || img.Flip == "H" || strings.ToLower(img.Flip) == "horizontal":
		flipObject.FlipH()
	case img.Flip == "v" || img.Flip == "V" || strings.ToLower(img.Flip) == "vertical":
		flipObject.FlipV()
	default:
		return flipObject.GetImage(), fmt.Errorf("we only support horizontal and vertical flips %s", img.Flip)
	}
	return flipObject.GetImage(), nil
}
