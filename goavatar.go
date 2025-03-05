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

const (
	gridSize = 8  // Grid size (8x8)
	scale    = 20 // Scale factor for pixel size
)

// get the any string and make a hash
func generateHash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// draw a single pixel and enlarge each pixle into 20x20 block
func drawPixel(img *image.RGBA, x, y int, c color.Color) {
	for dx := 0; dx < scale; dx++ {
		for dy := 0; dy < scale; dy++ {
			img.Set(x*scale+dx, y*scale+dy, c)
		}
	}
}

func Make(
	input,
	filename string,
) {
	hash := generateHash(input)

	// create a blank image
	img := image.NewRGBA(image.Rect(0, 0, gridSize*scale, gridSize*scale))

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
				drawPixel(img, x, y, avatarColor)
				drawPixel(img, gridSize-1-x, y, avatarColor) // mirror the pixel
			} else {
				drawPixel(img, x, y, bgColor)
				drawPixel(img, gridSize-1-x, y, bgColor) // mirror the bg pixel
			}

		}
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	png.Encode(file, img)
}
