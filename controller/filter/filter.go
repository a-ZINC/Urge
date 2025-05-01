package filter

import "image"

type Filter interface {
	Filter()
	GetOutputImage() image.Image
}