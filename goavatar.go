package goavatar

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/color"
)

// option contains the configuration for the avatar generator.
type options struct {
	size     int
	gridSize int
	bgColor  color.RGBA
	fgColor  color.RGBA
}

// optFunc is a function that applies an option to the options struct.
type OptFunc func(*options)

// WithSize sets the width and height of the avatar minimum 64x64.
func WithSize(s int) OptFunc {
	return func(o *options) {
		// insure that image should be at least 64x64
		if s >= 64 {
			o.size = s
		}
	}
}

// WithGridSize sets the grid size of the avatar.
func WithGridSize(g int) OptFunc {
	return func(o *options) {
		// make sure grid is minimum 8 to make nice pattrens
		if g > 8 {
			o.gridSize = g
		}
	}
}

// WithBgColor sets the background color of the avatar.
func WithBgColor(r, g, b, a uint8) OptFunc {
	return func(o *options) {
		o.bgColor = color.RGBA{r, g, b, a}
	}
}

// WithFgColor sets the foreground color of the avatar.
func WithFgColor(r, g, b, a uint8) OptFunc {
	return func(o *options) {
		o.fgColor = color.RGBA{r, g, b, a}
	}
}

// defaultOptions provides the default value to generate the avatar.
func defaultOptions(hash string) options {
	return options{
		size:     64,                                         // default size should be 64 to make sure images are perfect square
		gridSize: 8,                                          // minimum size for the grid for make shape complexity
		bgColor:  color.RGBA{240, 240, 240, 255},             // light gray color
		fgColor:  color.RGBA{hash[0], hash[1], hash[2], 255}, // use the first three hash bytes as the foreground color
	}
}

// generateHash generates the MD5 hash of the input string.
func generateHash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// drawPixel draws a single pixel and enlarge each pixle into 20x20 block.
func drawPixel(img *image.RGBA, x, y int, c color.Color, pixelW, pixelH int) {
	for dx := 0; dx < pixelW; dx++ {
		for dy := 0; dy < pixelH; dy++ {
			img.Set(x*pixelW+dx, y*pixelH+dy, c)
		}
	}
}

// Make generates an avatar image based on the input string and options.
func Make(input string, opts ...OptFunc) image.Image {
	// generate the hash of an input
	hash := generateHash(input)
	o := defaultOptions(hash)

	for _, opt := range opts {
		opt(&o)
	}

	// create a blank image
	img := image.NewRGBA(image.Rect(0, 0, o.size, o.size))

	pixelSizeX := o.size / o.gridSize // each grid cell width
	pixelSizeY := o.size / o.gridSize // each grid cell height

	// generate colors
	avatarColor := o.fgColor
	bgColor := o.bgColor

	// generate the pixel pattern
	// loop over each pixel in the grid
	for y := 0; y < o.gridSize; y++ {
		for x := 0; x < o.gridSize/2; x++ {
			// use bitwise operation to determine if a pixel should be colored
			pixelOn := (hash[y]>>(x%8))&1 == 1

			// image should
			if pixelOn {
				drawPixel(img, x, y, avatarColor, pixelSizeX, pixelSizeY)
				drawPixel(img, o.gridSize-1-x, y, avatarColor, pixelSizeX, pixelSizeY) // mirror the pixel
			} else {
				drawPixel(img, x, y, bgColor, pixelSizeX, pixelSizeY)
				drawPixel(img, o.gridSize-1-x, y, bgColor, pixelSizeX, pixelSizeY) // mirror the bg pixel
			}

		}
	}

	return img
}
