package rotate

import (
	"image"
	"urge/model"
)

type Rotate struct {
	img model.Image
}

func New(img model.Image) *Rotate {
	return &Rotate{img: img}
}

func (r *Rotate) GetImage() model.Image {
	return r.img
}

func (r *Rotate) Rotate90() {
	inputImage := r.img.Image
	if r.img.Output != nil {
		inputImage = r.img.Output
	}
	bounds := inputImage.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	newImage := image.NewRGBA(image.Rect(0, 0, height, width))
	for x := range width {
		for y := range height {
			newImage.Set(height-y-1, x, inputImage.At(x, y))
		}
	}
	r.img.Output = newImage
}

func (r *Rotate) Rotate180() {
	inputImage := r.img.Image
	if r.img.Output != nil {
		inputImage = r.img.Output
	}
	bounds := inputImage.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	newImage := image.NewRGBA(bounds)
	for x := range width {
		for y := range height {
			newImage.Set(width-x-1, height-y-1, inputImage.At(x, y))
		}
	}
	r.img.Output = newImage
}

func (r *Rotate) Rotate270() {
	inputImage := r.img.Image
	if r.img.Output != nil {
		inputImage = r.img.Output
	}
	bounds := inputImage.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	newImage := image.NewRGBA(image.Rect(0, 0, height, width))
	for x := range width {
		for y := range height {
			newImage.Set(y, x, inputImage.At(x, y))
		}
	}
	r.img.Output = newImage
}
