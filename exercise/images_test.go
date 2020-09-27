package exercise

import (
	"image"
	"image/color"
	"testing"

	"golang.org/x/tour/pic"
)

type Image struct {
	Height, Width int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Height, m.Width)
}

func (m Image) At(x, y int) color.Color {
	c := uint8(x ^ y)
	return color.RGBA{c, c, 255, 255}
}

func TestImages(t *testing.T) {
	m := Image{256, 256}
	pic.ShowImage(m)
}
