package goavatar

import (
	"image/color"
	"testing"
)

// expectedTopLeftPixel computes what color should appear at (0,0) by replaying the default options
// and then using the same raw hash logic as in Make: for x=0, y=0, it tests if (hash[0] & 1) == 1.
//
// NOTE: generateHash returns a hex‑encoded string, so here we use its first character’s ASCII code.
func expectedTopLeftPixel(input string, opts []optFunc) (col color.Color) {
	// generate the hash of the input
	hash := generateHash(input)
	// get the default configuration; which sets fgColor to {hash[0], hash[1], hash[2], 255}
	conf := defaultOptions(hash)
	// apply all option functions to the default configuration
	for _, opt := range opts {
		opt(&conf)
	}
	// For the top‐left cell (x=0,y=0), the decision is based on the least‐significant bit of the raw hash character.
	// Using the raw ASCII value of hash[0] as in the current implementation.
	if (hash[0] & 1) == 1 {
		return conf.fgColor
	}
	return conf.bgColor
}

func TestMake(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		opts   []optFunc
		width  int
		height int
	}{
		{
			name:  "Default settings",
			input: "test@example.com",
			opts:  nil, // defaults
			width: 256, height: 256,
		},
		{
			name:  "Custom width and height",
			input: "custom-size",
			opts:  []optFunc{WithWidth(512), WithHeight(512)},
			width: 512, height: 512,
		},
		{
			name:  "Custom background color",
			input: "custom-bg",
			// override background color only
			opts:  []optFunc{WithBgColor(255, 0, 0, 255)},
			width: 256, height: 256,
		},
		{
			name:  "Custom foreground color",
			input: "custom-fg",
			// override foreground color only
			opts:  []optFunc{WithFgColor(10, 20, 30, 255)},
			width: 256, height: 256,
		},
		{
			name:  "QuantumNomad42",
			input: "QuantumNomad42",
			opts:  []optFunc{WithWidth(512), WithHeight(512)},
			width: 512, height: 512,
		},
		{
			name:  "EchoFrost7",
			input: "EchoFrost7",
			opts:  []optFunc{WithWidth(512), WithHeight(512)},
			width: 512, height: 512,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			img := Make(tt.input, tt.opts...)
			if img == nil {
				t.Fatalf("Make() returned nil for input %q", tt.input)
			}

			// Verify the image dimensions.
			bounds := img.Bounds()
			if bounds.Dx() != tt.width || bounds.Dy() != tt.height {
				t.Errorf("Unexpected image size for %q: got %dx%d, want %dx%d",
					tt.input, bounds.Dx(), bounds.Dy(), tt.width, tt.height)
			}

			// Compute the expected top-left pixel color.
			expected := expectedTopLeftPixel(tt.input, tt.opts)
			actual := img.At(0, 0)
			ar, ag, ab, aa := actual.RGBA()
			er, eg, eb, ea := expected.RGBA()
			if ar != er || ag != eg || ab != eb || aa != ea {
				t.Errorf("Unexpected top-left pixel color for %q: got %v, want %v",
					tt.input, actual, expected)
			}
		})
	}
}
