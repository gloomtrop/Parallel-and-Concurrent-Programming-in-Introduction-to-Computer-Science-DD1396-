package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	image := make([][]uint8, dy)
	for y :=0; y < dy; y++{
		img_rows := make([]uint8, dx)
		for x:=0; x < dx; x++{
			img_rows[x] = uint8(x*10+y/3)
		}
		image[y] = img_rows
	} 
	return image
}

func main() {

	pic.Show(Pic)
}

