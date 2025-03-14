package goavatar

import (
	"image/color"
	"testing"
)

func TestMake(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		opts   Options
		width  int
		height int
	}{
		{
			name:   "Default settings",
			input:  "test@example.com",
			opts:   Options{}, // empty options for use the defaults
			width:  256,
			height: 256,
		},
		{
			name:  "Custom width and height",
			input: "custom-size",
			opts:  Options{Width: 512, Height: 512},
			width: 512, height: 512,
		},
		{
			name:  "Custom background color",
			input: "custom-bg",
			opts:  Options{BgColor: color.RGBA{255, 0, 0, 255}}, // Red background
			width: 256, height: 256,
		},
		{
			name:   "QuantumNomad42",
			input:  "QuantumNomad42",
			opts:   Options{Width: 512, Height: 512},
			width:  512,
			height: 512,
		},
		{
			name:   "EchoFrost7",
			input:  "EchoFrost7",
			opts:   Options{Width: 512, Height: 512},
			width:  512,
			height: 512,
		},
	}

	for _, tt := range tests {

		img := Make(tt.input, tt.opts)

		// check if the return image is not nill
		if img == nil {
			t.Fatalf("Make() returned nil for input %q", tt.input)
		}

		// check if dimensions is matched the expected size
		bounds := img.Bounds()
		if bounds.Dx() != tt.width || bounds.Dy() != tt.height {
			t.Errorf("Unexepected image size for %q: go %dx%d, want %dx%d",
				tt.input,
				bounds.Dx(),
				bounds.Dy(),
				tt.width,
				tt.height,
			)
		}

		// check if the top left pixel matches the BG color
		// check BgColor is set not an empty
		if tt.opts.BgColor != (color.RGBA{}) {
			expectedBgColor := tt.opts.BgColor
			actualColor := img.At(0, 0).(color.RGBA) // get the top left pixel color

			if actualColor != expectedBgColor {
				t.Errorf("Unexepected background color for %q: got %v, want %v", tt.input, actualColor, expectedBgColor)
			}

		}

	}
}
