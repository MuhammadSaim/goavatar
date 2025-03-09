package main

import (
	"image/color"

	"github.com/MuhammadSaim/goavatar"
)

func main() {
	// Generate the first avatar with a custom width and height
	options1 := goavatar.Options{
		Width:  512, // Set custom image width (default is 256)
		Height: 512, // Set custom image height (default is 256)
	}
	goavatar.Make("QuantumNomad42", "../arts/avatar_1.png", options1)
	// Generates a unique avatar based on "QuantumNomad42" and saves it as avatar_1.png

	// Generate the second avatar with a custom grid size
	options2 := goavatar.Options{
		Width:    512, // Set custom image width (default is 256)
		Height:   512, // Set custom image height (default is 256)
		GridSize: 10,  // Set custom grid size (default is 8), affects pattern complexity
	}
	goavatar.Make("EchoFrost7", "../arts/avatar_2.png", options2)
	// Generates an avatar with a 10x10 grid for more detail and saves it as avatar_2.png

	// Generate the third avatar with a custom background color
	options3 := goavatar.Options{
		Width:   512,                           // Set custom image width (default is 256)
		Height:  512,                           // Set custom image height (default is 256)
		BgColor: color.RGBA{170, 120, 10, 255}, // Change background color (default is light gray)
	}
	goavatar.Make("NebulaTide19", "../arts/avatar_3.png", options3)
	// Generates an avatar with a brownish background color and saves it as avatar_3.png

	// Generate the fourth avatar with a custom foreground and background color
	options4 := goavatar.Options{
		Width:   512,                            // Set custom image width (default is 256)
		Height:  512,                            // Set custom image height (default is 256)
		BgColor: color.RGBA{170, 120, 10, 255},  // Change background color (default is light gray)
		FgColor: color.RGBA{255, 255, 255, 255}, // Change foreground color (default is extracted from hash)
	}
	goavatar.Make("ZephyrPulse88", "../arts/avatar_4.png", options4)
	// Generates an avatar with a brownish background and white foreground, saving it as avatar_4.png

	// Generate an avatar using default settings
	goavatar.Make("EmberNexus23", "../arts/avatar_5.png", goavatar.Options{})
	// Uses default width (256), height (256), grid size (8), and colors
	// Saves the generated avatar as avatar_5.png
}
