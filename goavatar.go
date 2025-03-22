package goavatar

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/color"
)

// option contains the configuration for the avatar generator.
type options struct {
	width    int
	height   int
	gridSize int
	bgColor  color.RGBA
	fgColor  color.RGBA
}

// optFunc is a function that applies an option to the options struct.
type optFunc func(*options)

// Width sets the width of the avatar.
func Width(w int) optFunc {
	return func(o *options) {
		o.width = w
	}
}

// Height sets the height of the avatar.
func Height(h int) optFunc {
	return func(o *options) {
		o.height = h
	}
}

// GridSize sets the grid size of the avatar.
func GridSize(g int) optFunc {
	return func(o *options) {
		o.gridSize = g
	}
}

// BgColor sets the background color of the avatar.
func BgColor(r, g, b, a uint8) optFunc {
	return func(o *options) {
		o.bgColor = color.RGBA{r, g, b, a}
	}
}

// FgColor sets the foreground color of the avatar.
func FgColor(r, g, b, a uint8) optFunc {
	return func(o *options) {
		o.fgColor = color.RGBA{r, g, b, a}
	}
}

// defaultOptions provides the default value to generate the avatar.
func defaultOptions(hash string) options {
	return options{
		width:    256,
		height:   256,
		gridSize: 8,
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
func Make(input string, opts ...optFunc) image.Image {
	// generate the hash of an input
	hash := generateHash(input)
	o := defaultOptions(hash)

	for _, opt := range opts {
		opt(&o)
	}

	// create a blank image
	img := image.NewRGBA(image.Rect(0, 0, o.width, o.height))

	pixelSizeX := o.width / o.gridSize  // each grid cell width
	pixelSizeY := o.height / o.gridSize // each grid cell height

	// generate colors
	avatarColor := o.fgColor
	bgColor := o.bgColor

	// generate the pixel pattern
	// loop over each pixel in the grid
	for y := 0; y < o.gridSize; y++ {
		for x := 0; x < o.gridSize/2+1; x++ {
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
