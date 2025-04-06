package goavatar

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/color"
	"math"
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

// drawPixel draws a single pixel block based on proportional scaling to avoid gaps.
func drawPixel(img *image.RGBA, gridX, gridY int, c color.Color, gridSize, imageSize int) {
	// Calculate exact scaled bounds
	startX := int(math.Round(float64(gridX) * float64(imageSize) / float64(gridSize)))
	startY := int(math.Round(float64(gridY) * float64(imageSize) / float64(gridSize)))
	endX := int(math.Round(float64(gridX+1) * float64(imageSize) / float64(gridSize)))
	endY := int(math.Round(float64(gridY+1) * float64(imageSize) / float64(gridSize)))

	// Clamp to image size to avoid out-of-bounds
	if endX > img.Bounds().Dx() {
		endX = img.Bounds().Dx()
	}
	if endY > img.Bounds().Dy() {
		endY = img.Bounds().Dy()
	}

	// Fill the block
	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			img.Set(x, y, c)
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

	// generate colors
	avatarColor := o.fgColor
	bgColor := o.bgColor
	isOdd := o.gridSize%2 != 0

	// generate the pixel pattern
	// loop over each pixel in the grid
	for y := 0; y < o.gridSize; y++ {
		for x := 0; x < o.gridSize/2; x++ {
			// use bitwise operation to determine if a pixel should be colored
			pixelOn := (hash[y]>>(x%8))&1 == 1

			// image should
			if pixelOn {
				drawPixel(img, x, y, avatarColor, o.gridSize, o.size)
				drawPixel(img, o.gridSize-1-x, y, avatarColor, o.gridSize, o.size) // mirror the pixel
			} else {
				drawPixel(img, x, y, bgColor, o.gridSize, o.size)
				drawPixel(img, o.gridSize-1-x, y, bgColor, o.gridSize, o.size) // mirror the bg pixel
			}

		}
		// Draw the center column if gridSize is odd
		if isOdd {
			mid := o.gridSize / 2
			pixelOn := (hash[y]>>(mid%8))&1 == 1
			color := bgColor
			if pixelOn {
				color = avatarColor
			}
			drawPixel(img, mid, y, color, o.gridSize, o.size)
		}
	}

	return img
}
