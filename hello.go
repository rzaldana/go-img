package main

import (
	"fmt"
	"image"
	"image/png"
	"image/color"
	"os"
)

func main() {
	fmt.Println("Hello again")
	
	// Create in-memory image
	height := 200 
	width := 100
	myImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// Define color red
	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	// Fill in the image with the red color
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			myImg.Set(x, y, red)
		}
	}

	// Crete image file
	out, _ := os.Create("image.png")

	// store in-memory image in file
	png.Encode(out, myImg)
	out.Close()
}
