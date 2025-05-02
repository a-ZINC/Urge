package utils

import (
	"urge/controller/filter"
	"urge/log"
	"urge/model"
)

func Filter(img model.Image) filter.Filter {
	log.InfoLogger.Printf("----------Filtering image------------")
	var filterType filter.Filter
	switch img.Filter {
	case "grayscale":
		filterType = filter.New(img)
	}
	if filterType != nil {
		filterType.Filter()
	}
	return filterType
}
