package filter

import "urge/model"

type Filter interface {
	Filter()
	GetImage() model.Image
}