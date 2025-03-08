package main

import (
	"os"

	"github.com/MuhammadSaim/goavatar"
)

func main() {
	file1, _ := os.Create("arts/avatar_1.png")
	defer file1.Close()
	goavatar.Make("QuantumNomad42", file1)

	file2, _ := os.Create("arts/avatar_2.png")
	defer file2.Close()
	goavatar.Make("EchoFrost7", file2)

	file3, _ := os.Create("arts/avatar_3.png")
	defer file3.Close()
	goavatar.Make("NebulaTide19", file3)

	file4, _ := os.Create("arts/avatar_4.png")
	defer file4.Close()
	goavatar.Make("ZephyrPulse88", file4)

	file5, _ := os.Create("arts/avatar_5.png")
	defer file5.Close()
	goavatar.Make("EmberNexus23", file5)
}
