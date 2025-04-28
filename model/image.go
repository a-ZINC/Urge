package model

import "image"

type Image struct {
	Image image.Image
	Resize string
	Filter string
	Url string
	Output string
	Format string
}