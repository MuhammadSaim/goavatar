package goavatar

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/color"
	"image/draw"
	"math"
)

// options contains the configuration for the avatar generator.
type options struct {
	size     int
	gridSize int
	bgColor  color.NRGBA
	fgColors []color.NRGBA
	layers   int 
}

// OptFunc is a function that applies an option to the options struct.
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
		o.bgColor = color.NRGBA{r, g, b, a}
	}
}

// WithFgColor sets the foreground color of the avatar.
// It sets the first layer's color.
func WithFgColor(r, g, b, a uint8) OptFunc {
	return func(o *options) {
		o.fgColors = []color.NRGBA{{r, g, b, a}}
	}
}

// WithLayers sets the number of foreground layers (1-3).
func WithLayers(n int) OptFunc {
	return func(o *options) {
		if n >= 1 && n <= 3 {
			o.layers = n
		}
	}
}

// WithLayerColor sets the color for a specific layer index (0-based).
func WithLayerColor(layerIndex int, r, g, b, a uint8) OptFunc {
	return func(o *options) {
		// Expand slice if needed
		for len(o.fgColors) <= layerIndex {
			o.fgColors = append(o.fgColors, color.NRGBA{})
		}
		o.fgColors[layerIndex] = color.NRGBA{r, g, b, a}
	}
}

// defaultOptions provides the default value to generate the avatar.
func defaultOptions(hash string) options {
	return options{
		size:     64,                            // default size should be 64 to make sure images are perfect square
		gridSize: 8,                             // minimum size for the grid for make shape complexity
		bgColor:  color.NRGBA{240, 240, 240, 255}, // light gray color
		fgColors: []color.NRGBA{{hash[0], hash[1], hash[2], 255}}, // use the first three hash bytes as the foreground color
		layers:   1,
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
	rect := image.Rect(startX, startY, endX, endY)
	draw.Draw(img, rect, &image.Uniform{c}, image.Point{}, draw.Over)
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

	// Fill background
	draw.Draw(img, img.Bounds(), &image.Uniform{o.bgColor}, image.Point{}, draw.Src)

	currentHash := hash
	isOdd := o.gridSize%2 != 0

	for l := 0; l < o.layers; l++ {
		// derive hash for this layer
		if l > 0 {
			currentHash = generateHash(currentHash)
		}

		// determine color
		var avatarColor color.NRGBA
		if l < len(o.fgColors) {
			avatarColor = o.fgColors[l]
			// Check if color is empty/zero? defaultOptions sets index 0.
			// WithLayerColor might extend with zeros.
			// If alpha is 0, should we generate? 
			// User might purposefully set transparent? Unlikely for avatar foreground.
			// Assuming if user sets it, they set it.
			// But if we expanded with empty RGBA (0,0,0,0), it's invisible.
			// If it is strictly 0,0,0,0, maybe fallback to hash? 
			// Let's assume user provides valid colors if they use WithLayerColor.
			// But for "unspecified" layers where user requested 3 layers but provided 1 color:
			if avatarColor == (color.NRGBA{}) {
				avatarColor = color.NRGBA{currentHash[0], currentHash[1], currentHash[2], 255}
			}
		} else {
			avatarColor = color.NRGBA{currentHash[0], currentHash[1], currentHash[2], 255}
		}

		// generate the pixel pattern
		// loop over each pixel in the grid
		for y := 0; y < o.gridSize; y++ {
			for x := 0; x < o.gridSize/2; x++ {
				// use bitwise operation to determine if a pixel should be colored
				pixelOn := (currentHash[y]>>(x%8))&1 == 1

				if pixelOn {
					drawPixel(img, x, y, avatarColor, o.gridSize, o.size)
					drawPixel(img, o.gridSize-1-x, y, avatarColor, o.gridSize, o.size) // mirror the pixel
				}
			}
			// Draw the center column if gridSize is odd
			if isOdd {
				mid := o.gridSize / 2
				pixelOn := (currentHash[y]>>(mid%8))&1 == 1
				if pixelOn {
					drawPixel(img, mid, y, avatarColor, o.gridSize, o.size)
				}
			}
		}
	}

	return img
}
