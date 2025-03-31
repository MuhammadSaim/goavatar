package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/MuhammadSaim/goavatar"
)

// smallShape and triangleShape define custom pixel shapes
// for use in some of the examples below
var smallShape = [][]int{
	{0, 0, 0, 0, 0},
	{0, 1, 1, 1, 0},
	{0, 1, 1, 1, 0},
	{0, 1, 1, 1, 0},
	{0, 0, 0, 0, 0},
}
var triangleShape = [][]int{
	{0, 0, 0, 0, 1},
	{0, 0, 0, 1, 1},
	{0, 0, 1, 1, 1},
	{0, 1, 1, 1, 1},
	{1, 1, 1, 1, 1},
}

// customSizeExample creates an avatar for "QuantumNomad42"
// with a custom image size (512)
func customSizeExample() image.Image {
	return goavatar.Make("QuantumNomad42",
		goavatar.WithSize(512), // Set custom image size (default is 64)
	)
}

// customSizeAndGridExample creates an avatar for "EchoFrost7"
// with a custom image size (512) and a custom grid size (10)
func customSizeAndGridExample() image.Image {
	return goavatar.Make("EchoFrost7",
		goavatar.WithSize(512),    // Set custom image size (default is 64)
		goavatar.WithGridSize(10), // Set custom grid size (default is 8), affects pattern complexity
	)
}

// customSizeAndBackgroundColorExample creates an avatar for "NebulaTide19"
// with a custom image size (100) and background color
func customSizeAndBackgroundColorExample() image.Image {
	return goavatar.Make("NebulaTide19",
		goavatar.WithSize(100),                  // Set custom image widthxheight (default is 64)
		goavatar.WithBgColor(170, 120, 10, 255), // Change background color (default is light gray)
	)
}

// customSizeAndColorsExample creates an avatar for "ZephyrPulse88"
// witha custom size (50) and foreground/background colors
func customSizeAndColorsExample() image.Image {
	return goavatar.Make("ZephyrPulse88",
		goavatar.WithSize(50),                    // Set custom image widthxheight if size is less then 64 this will go to default (default is 64)
		goavatar.WithBgColor(170, 120, 10, 255),  // Change background color (default is light gray)
		goavatar.WithFgColor(255, 255, 255, 255), // Change foreground color (default is extracted from hash)
	)
}

// defaultExample() creates an avatar for "EmberNexus23" with the default settings
func defaultExample() image.Image {
	return goavatar.Make("EmberNexus23")
}

// customForegroundShapeExample() creates an avatar for "EmberNexus23" with
// a custom-shaped foreground pixel
func customForegroundShapeExample() image.Image {
	return goavatar.Make("EmberNexus23",
		goavatar.WithSize(512),    // Set custom image size (default is 64)
		goavatar.WithGridSize(10), // Set custom grid size (default is 8), affects pattern complexity
		goavatar.WithFgShape(smallShape),
	)
}

// customPixelShapesExample() creates an avatar for "EmberNexus23" with
// custom-shaped foreground and background pixels
func customPixelShapesExample() image.Image {
	return goavatar.Make("EmberNexus23",
		goavatar.WithSize(512),    // Set custom image size (default is 64)
		goavatar.WithGridSize(10), // Set custom grid size (default is 8), affects pattern complexity
		goavatar.WithFgShape(smallShape),
		goavatar.WithBgShape(triangleShape),
	)
}

// customTransparentForegroundShapeExample() creates an avatar for "EmberNexus23" with
// a custom-shaped foreground pixel
// and transparency to allow the background to show through gaps in foreground pixels
func customTransparentForegroundShapeExample() image.Image {
	return goavatar.Make("EmberNexus23",
		goavatar.WithSize(512),    // Set custom image size (default is 64)
		goavatar.WithGridSize(10), // Set custom grid size (default is 8), affects pattern complexity
		goavatar.WithFgShape(smallShape),
		goavatar.WithTransparency(),
	)
}

// customTransparentPixelShapesExample() creates an avatar for "EmberNexus23" with
// custom-shaped foreground and background pixels
// and transparency to allow the background to show through gaps in foreground pixels
func customTransparentPixelShapesExample() image.Image {
	return goavatar.Make("EmberNexus23",
		goavatar.WithSize(512),    // Set custom image size (default is 64)
		goavatar.WithGridSize(10), // Set custom grid size (default is 8), affects pattern complexity
		goavatar.WithFgShape(smallShape),
		goavatar.WithBgShape(triangleShape),
		goavatar.WithTransparency(),
	)
}

// createFileForImage creates an example file for a given image.
// The filename is created based on the image number argument.
func createFileForImage(n int, img image.Image) {
	fn := fmt.Sprintf("../arts/avatar_%d.png", n)

	file, err := os.Create(fn)
	if err != nil {
		fmt.Printf("Error creating file %s: %s\n", fn, err)
		return
	}
	defer file.Close()

	// Encode image as PNG and save
	err = png.Encode(file, img)
	if err != nil {
		fmt.Printf("Error saving image %d: %s\n", n, err)
		return
	}
	fmt.Println("Saved: ", fn)
}

func main() {
	createFileForImage(1, customSizeExample())
	createFileForImage(2, customSizeAndGridExample())
	createFileForImage(3, customSizeAndBackgroundColorExample())
	createFileForImage(4, customSizeAndColorsExample())
	createFileForImage(5, defaultExample())
	createFileForImage(6, customForegroundShapeExample())
	createFileForImage(7, customTransparentForegroundShapeExample())
	createFileForImage(8, customPixelShapesExample())
	createFileForImage(9, customTransparentPixelShapesExample())
}
