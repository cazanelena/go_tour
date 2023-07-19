package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Create a 2D slice to store the image
	picture := make([][]uint8, dy)

	// Loop to allocate each []uint8 slice
	for y := 0; y < dy; y++ {
		picture[y] = make([]uint8, dx)

		// Loop to calculate the value for each pixel
		for x := 0; x < dx; x++ {
			// Calculate intensity using the function (x+y)/2
			intensity := x*y
			// Convert the intensity to a uint8 value and store it in the slice
			picture[y][x] = uint8(intensity)
		}
	}

	return picture
}

func main() {
	pic.Show(Pic)
}
