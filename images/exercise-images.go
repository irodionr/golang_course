package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image - implementation of image.Image
type Image struct {
	Rect image.Rectangle
}

// ColorModel - return RGBAModel
func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds - return Rectangle
func (img Image) Bounds() image.Rectangle {
	return img.Rect
}

// At - return color
func (Image) At(x, y int) color.Color {
	return color.RGBA{uint8((x + y) / 2), uint8((x + y) / 2), 255, 255}
}

func main() {
	m := Image{image.Rect(0, 0, 256, 256)}
	pic.ShowImage(m) // Base64 to PNG: https://codebeautify.org/base64-to-image-converter
}
