package goavatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Option struct for the configurations
type Options struct {
	Width    int
	Height   int
	GridSize int
	BgColor  color.RGBA
	FgColor  color.RGBA
}

// default options to provide the default value
var defaultOptions = Options{
	Width:    256,
	Height:   256,
	GridSize: 8,
	BgColor:  color.RGBA{240, 240, 240, 255}, // light gray color
}

// get the any string and make a hash
func generateHash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// draw a single pixel and enlarge each pixle into 20x20 block
func drawPixel(img *image.RGBA, x, y int, c color.Color, pixelW, pixelH int) {
	for dx := 0; dx < pixelW; dx++ {
		for dy := 0; dy < pixelH; dy++ {
			img.Set(x*pixelW+dx, y*pixelH+dy, c)
		}
	}
}

func Make(
	input,
	filename string,
	opts Options,
) {
	// generate the hash of an input
	hash := generateHash(input)

	// use default options if any value is empty or 0
	if opts.Width == 0 || opts.Width < 100 {
		opts.Width = defaultOptions.Width
	}

	if opts.Height == 0 || opts.Height < 100 {
		opts.Height = defaultOptions.Height
	}

	if opts.GridSize == 0 || opts.GridSize < 8 {
		opts.GridSize = defaultOptions.GridSize
	}

	if opts.BgColor == (color.RGBA{}) {
		opts.BgColor = defaultOptions.BgColor
	}

	if opts.FgColor == (color.RGBA{}) {
		opts.FgColor = color.RGBA{hash[0], hash[1], hash[2], 255}
	}

	// create a blank image
	img := image.NewRGBA(image.Rect(0, 0, opts.Width, opts.Height))

	pixelSizeX := opts.Width / opts.GridSize  // each grid cell width
	pixelSizeY := opts.Height / opts.GridSize // each grid cell height

	// generate colors
	avatarColor := opts.FgColor
	bgColor := opts.BgColor

	// generate the pixel patren
	// loop over each pixel in the grid
	for y := 0; y < opts.GridSize; y++ {
		for x := 0; x < opts.GridSize/2+1; x++ {
			// use bitwise operation to determine if a pixel should be colored
			pixelOn := (hash[y]>>(x%8))&1 == 1

			// image should
			if pixelOn {
				drawPixel(img, x, y, avatarColor, pixelSizeX, pixelSizeY)
				drawPixel(img, opts.GridSize-1-x, y, avatarColor, pixelSizeX, pixelSizeY) // mirror the pixel
			} else {
				drawPixel(img, x, y, bgColor, pixelSizeX, pixelSizeY)
				drawPixel(img, opts.GridSize-1-x, y, bgColor, pixelSizeX, pixelSizeY) // mirror the bg pixel
			}

		}
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	png.Encode(file, img)

	err = file.Sync()
	if err != nil {
		fmt.Println("Error syncing file: ", err)
	}

	file.Close()
}
