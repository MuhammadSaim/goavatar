package goavatar

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/color"
)

// option contains the configuration for the avatar generator.
type options struct {
	size         int
	gridSize     int
	bgColor      color.RGBA
	fgColor      color.RGBA
	fgShape      [][]int
	bgShape      [][]int
	transparency bool
}

// optFunc is a function that applies an option to the options struct.
type optFunc func(*options)

// WithSize sets the size of the avatar.
// It is always square and has a minimum size of 64x64
func WithSize(s int) optFunc {
	return func(o *options) {
		// insure that image should be at least 64x64
		if s >= 64 {
			o.size = s
		}
	}
}

// WithGridSize sets the grid size of the avatar.
func WithGridSize(g int) optFunc {
	return func(o *options) {
		o.gridSize = g
	}
}

// WithBgColor sets the background color of the avatar.
func WithBgColor(r, g, b, a uint8) optFunc {
	return func(o *options) {
		o.bgColor = color.RGBA{r, g, b, a}
	}
}

// WithFgColor sets the foreground color of the avatar.
func WithFgColor(r, g, b, a uint8) optFunc {
	return func(o *options) {
		o.fgColor = color.RGBA{r, g, b, a}
	}
}

// WithFgShape sets the shape of the foreground pixels.
func WithFgShape(s [][]int) optFunc {
	return func(o *options) {
		o.fgShape = s
	}
}

// WithBgShape sets the shape of the background pixels.
func WithBgShape(s [][]int) optFunc {
	return func(o *options) {
		o.bgShape = s
	}
}

// WithTransparency sets the option to have the background show
// through any negative space in shaped foreground pixels.
func WithTransparency() optFunc {
	return func(o *options) {
		o.transparency = true
	}
}

// defaultOptions provides the default value to generate the avatar.
func defaultOptions(hash string) options {
	return options{
		size:     64, // default size should be 64 to make sure images are perfect square
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
func drawPixel(img *image.RGBA, x, y int, c color.Color, size int) {
	for dx := 0; dx < size; dx++ {
		for dy := 0; dy < size; dy++ {
			img.Set(x*size+dx, y*size+dy, c)
		}
	}
}

// drawShapedPixel draws a single logical pixel that matches the supplied bitmap
func drawShapedPixel(img *image.RGBA, x, y int, c color.Color, shape [][]int) {
	height := len(shape)
	if height == 0 {
		return
	}
	width := len(shape[0])

	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			if shape[dy][dx] != 0 {
				img.Set(x*width+dx, y*height+dy, c)
			}
		}
	}
}

// scaleShape scales up a bitmap of what each shaped pixel should look like
// returns nil if shape is nil
func scaleShape(shape [][]int, targetW int) [][]int {
	if shape == nil {
		return nil
	}

	origH := len(shape)
	if origH == 0 {
		return nil
	}
	origW := len(shape[0])

	scale := float64(targetW) / float64(origW)
	targetH := int(float64(origH) * scale)

	// Create scaled shape
	scaled := make([][]int, targetH)
	for y := 0; y < targetH; y++ {
		scaled[y] = make([]int, targetW)
		for x := 0; x < targetW; x++ {
			origY := int(float64(y) / scale)
			origX := int(float64(x) / scale)
			scaled[y][x] = shape[origY][origX]
		}
	}
	return scaled
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
	img := image.NewRGBA(image.Rect(0, 0, o.size, o.size))

	pixelSize := o.size / o.gridSize // each grid cell size

	// generate colors
	avatarColor := o.fgColor
	bgColor := o.bgColor

	// scale shapes to their correct size
	bgShape := scaleShape(o.bgShape, pixelSize)
	fgShape := scaleShape(o.fgShape, pixelSize)

	// generate the pixel pattern
	// loop over each pixel in the grid
	for y := 0; y < o.gridSize; y++ {
		for x := 0; x < o.gridSize/2; x++ {
			// use bitwise operation to determine if a pixel should be colored
			pixelOn := (hash[y]>>(x%8))&1 == 1

			// if this is a background pixel *or* it's foreground but we want the background to show through
			if !pixelOn || o.transparency {

				if bgShape != nil {
					drawShapedPixel(img, x, y, bgColor, bgShape)              // draw a shaped pixel
					drawShapedPixel(img, o.gridSize-1-x, y, bgColor, bgShape) // mirror the shaped pixel
				} else {
					drawPixel(img, x, y, bgColor, pixelSize)              // draw a normal pixel
					drawPixel(img, o.gridSize-1-x, y, bgColor, pixelSize) // mirror the pixel
				}
			}

			// if this is a foreground pixel
			if pixelOn {
				if fgShape != nil {
					drawShapedPixel(img, x, y, avatarColor, fgShape)
					drawShapedPixel(img, o.gridSize-1-x, y, avatarColor, fgShape) // mirror the shaped pixel
				} else {
					drawPixel(img, x, y, avatarColor, pixelSize)
					drawPixel(img, o.gridSize-1-x, y, avatarColor, pixelSize) // mirror the pixel
				}
			}

		}
	}

	return img
}
