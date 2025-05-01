package filter

import (
	"fmt"
	"image"
	"image/color"
)

type GrayScale struct {
	img image.Image
	Out image.Image
}

func New(img image.Image) *GrayScale {
	return &GrayScale{
		img: img,
	}
}

func (g *GrayScale) Filter() {
	fmt.Println("grayscale", g.img.ColorModel().Convert(color.Black))
	bounds := g.img.Bounds()
	grayImage := image.NewGray(bounds)

	for y:= bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := g.img.At(x, y).RGBA()
			gray := (0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			grayImage.SetGray(x, y, color.Gray{uint8(gray)})
		}
	}
	g.Out = grayImage
}

func (g *GrayScale) GetOutputImage() image.Image {
	return g.Out
}
