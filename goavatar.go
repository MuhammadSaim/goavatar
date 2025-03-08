package goavatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
)

const (
	width    = 256
	height   = 256
	gridSize = 8 // Grid size (8x8)
)

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
	input string,
	w io.Writer,
) {
	hash := generateHash(input)

	// create a blank image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	pixelSizeX := width / gridSize  // each grid cell width
	pixelSizeY := height / gridSize // each grid cell height

	// generate colors
	avatarColor := color.RGBA{hash[0], hash[1], hash[2], 255}
	bgColor := color.RGBA{240, 240, 240, 255}

	// generate the pixel patren
	// loop over each pixel in the grid
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize/2+1; x++ {
			// use bitwise operation to determine if a pixel should be colored
			pixelOn := (hash[y]>>(x%8))&1 == 1

			// image should
			if pixelOn {
				drawPixel(img, x, y, avatarColor, pixelSizeX, pixelSizeY)
				drawPixel(img, gridSize-1-x, y, avatarColor, pixelSizeX, pixelSizeY) // mirror the pixel
			} else {
				drawPixel(img, x, y, bgColor, pixelSizeX, pixelSizeY)
				drawPixel(img, gridSize-1-x, y, bgColor, pixelSizeX, pixelSizeY) // mirror the bg pixel
			}

		}
	}

	if err := png.Encode(w, img); err != nil {
		fmt.Printf("Error encoding image %s: %v\n", input, err)
		return
	}
}
