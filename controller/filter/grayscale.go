package filter

import (
	"fmt"
	"image"
)

type GrayScale struct {
	img image.Image
}

func New(img image.Image) *GrayScale {
	return &GrayScale{
		img: img,
	}
}

func (g *GrayScale) Filter() {
	fmt.Println(g.img)
}
