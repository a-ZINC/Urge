package utils

import (
	"fmt"
	"urge/controller/filter"
	"urge/log"
	"urge/model"
)

func Filter(img model.Image) (filter.Filter, error) {
	log.InfoLogger.Printf("----------Filtering image------------")
	var filterType filter.Filter
	switch img.Filter {
	case "grayscale":
		filterType = filter.New(img)
	default:
		return filterType, fmt.Errorf("we only support grayscale filter")
	}
	if filterType != nil {
		filterType.Filter()
	}
	return filterType, nil
}
