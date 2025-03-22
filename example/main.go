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
		goavatar.Width(512),  // Set custom image width (default is 256)
		goavatar.Height(512), // Set custom image height (default is 256)
	)

	// Generate the second avatar with a custom grid size with a 10x10 grid for more detail.
	// Saves the generated avatar as avatar_2.png
	image2 := goavatar.Make("EchoFrost7",
		goavatar.Width(512),   // Set custom image width (default is 256)
		goavatar.Height(512),  // Set custom image height (default is 256)
		goavatar.GridSize(10), // Set custom grid size (default is 8), affects pattern complexity
	)

	// Generate the third avatar with a custom brownish background color.
	// Saves the generated avatar as avatar_3.png
	image3 := goavatar.Make("NebulaTide19",
		goavatar.Width(512),                 // Set custom image width (default is 256)
		goavatar.Height(512),                // Set custom image height (default is 256)
		goavatar.BgColor(170, 120, 10, 255), // Change background color (default is light gray)
	)

	// Generate the fourth avatar with a custom brownish background and white foreground.
	// Saves the generated avatar as avatar_4.png
	image4 := goavatar.Make("ZephyrPulse88",
		goavatar.Width(512),                  // Set custom image width (default is 256)
		goavatar.Height(512),                 // Set custom image height (default is 256)
		goavatar.BgColor(170, 120, 10, 255),  // Change background color (default is light gray)
		goavatar.FgColor(255, 255, 255, 255), // Change foreground color (default is extracted from hash)

	)

	// Generate an avatar using default settings
	// Saves the generated avatar as avatar_5.png
	image5 := goavatar.Make("EmberNexus23")

	// append all the images into the list
	imgSlice = append(imgSlice, image1, image2, image3, image4, image5)

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
