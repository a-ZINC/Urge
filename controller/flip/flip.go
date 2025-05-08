package flip

import (
	"image"
	"urge/model"
)

type Flip struct {
	img model.Image
}

func New(img model.Image) *Flip {
	return &Flip{img: img}
}

func (r *Flip) GetImage() *model.Image {
	return &r.img
}

func (f *Flip) FlipH() {
	bounds := f.img.Image.Bounds()
	flippedImage := image.NewRGBA(bounds)
	for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
		for i := bounds.Min.X; i < bounds.Max.X; i++ {
			flippedImage.Set(bounds.Max.X - i, j, f.img.Image.At(i, j))
		}
	}
	f.img.Output = flippedImage
}

func (f *Flip) FlipV() {
	bounds := f.img.Image.Bounds()
	flippedImage := image.NewRGBA(bounds)
	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			flippedImage.Set(i, bounds.Max.Y - j, f.img.Image.At(i, j))
		}
	}
	f.img.Output = flippedImage
}
