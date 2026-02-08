package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/MuhammadSaim/goavatar"
)

func main() {
	// empty slice.
	imgSlice := make([]image.Image, 0)

	// Generates a unique avatar based on "QuantumNomad42" with a custom width and height.
	// Saves the generated avatar as avatar_1.png
	image1 := goavatar.Make("QuantumNomad42",
		goavatar.WithSize(512), // Set custom image widthxheight (default is 64)
	)

	// Generate the second avatar with a custom grid size with a 10x10 grid for more detail.
	// Saves the generated avatar as avatar_2.png
	image2 := goavatar.Make("EchoFrost7",
		goavatar.WithSize(512),    // Set custom image widthxheight (default is 64)
		goavatar.WithGridSize(10), // Set custom grid size (default is 8), affects pattern complexity
	)

	// Generate the third avatar with a custom brownish background color.
	// Saves the generated avatar as avatar_3.png
	image3 := goavatar.Make("NebulaTide19",
		goavatar.WithSize(100),                  // Set custom image widthxheight (default is 64)
		goavatar.WithBgColor(170, 120, 10, 255), // Change background color (default is light gray)
	)

	// Generate the fourth avatar with a custom brownish background and white foreground.
	// Saves the generated avatar as avatar_4.png
	image4 := goavatar.Make("ZephyrPulse88",
		goavatar.WithSize(50),                    // Set custom image widthxheight if size is less then 64 this will go to default (default is 64)
		goavatar.WithBgColor(170, 120, 10, 255),  // Change background color (default is light gray)
		goavatar.WithFgColor(255, 255, 255, 255), // Change foreground color (default is extracted from hash)
	)

	// Generate an avatar using default settings
	// Saves the generated avatar as avatar_5.png
	image5 := goavatar.Make("EmberNexus23")

	// Collect options dynamically
	var opts []goavatar.OptFunc

	// add size
	opts = append(opts, goavatar.WithSize(500))
	opts = append(opts, goavatar.WithGridSize(13))
	image6 := goavatar.Make("nice__user__name", opts...)

	// Generate multi-layered avatar (2 layers)
	image7 := goavatar.Make("MultiLayer2",
		goavatar.WithSize(512),
		goavatar.WithLayers(2),
	)

	// Generate multi-layered avatar (3 layers) with custom colors for each layer
	image8 := goavatar.Make("MultiLayer3Custom",
		goavatar.WithSize(512),
		goavatar.WithLayers(3),
		goavatar.WithLayerColor(0, 255, 0, 0, 255),   // Red
		goavatar.WithLayerColor(1, 0, 255, 0, 255),   // Green
		goavatar.WithLayerColor(2, 0, 0, 255, 255),   // Blue
	)

	// append all the images into the list
	imgSlice = append(imgSlice, image1, image2, image3, image4, image5, image6, image7, image8)

	// loop through the image slice and save the images
	for i, img := range imgSlice {

		filename := fmt.Sprintf("../arts/avatar_%d.png", i+1)

		// Create the file
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer file.Close()

		// Encode image as PNG and save
		err = png.Encode(file, img)
		if err != nil {
			fmt.Println("Error saving image:", err)
		} else {
			fmt.Println("Saved: ", filename)
		}

	}
}
