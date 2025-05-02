package filter

import (
	"fmt"
	"image"
	"image/color"
	"urge/model"
)

type GrayScale struct {
	img model.Image
}

func New(img model.Image) *GrayScale {
	return &GrayScale{
		img: img,
	}
}

func (g *GrayScale) Filter() {
	fmt.Println("grayscale", g.img.Image.ColorModel().Convert(color.Black))
	inputImage := g.img.Image
	if g.img.Output != nil {
		inputImage = g.img.Output
	}
	bounds := inputImage.Bounds()
	grayImage := image.NewGray(bounds)

	for y:= bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := inputImage.At(x, y).RGBA()
			gray := (0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			grayImage.SetGray(x, y, color.Gray{uint8(gray)})
		}
	}
	g.img.Output = grayImage
}

func (g *GrayScale) GetImage() model.Image {
	return g.img
}
