package model

import "image"

type Image struct {
	Image image.Image
	Resize string
	Filter string
	Rotate int
	Url string
	OutputUrl string
	Format string
	Output image.Image
}