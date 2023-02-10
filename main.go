package main

import (
	"fmt"
	"golang-ascii/img"
	"os"
)

const density = "Ñ@#W$9876543210?!abc;:+=-,._ "
const max = 765

func main() {
	img, err := img.Load()

	if err != nil {
		panic(err)
	}

	f, err := os.Create("./result.txt")

	if err != nil {
		panic(err)
	}

	err = os.Truncate("./result.txt", 0)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i := 0; i < img.Bounds().Max.Y-1; i++ {
		for j := 0; j < img.Bounds().Max.X-1; j++ {
			r8, g8, b8, _ := img.At(j, i).RGBA()

			r := r8 >> 8
			g := g8 >> 8
			b := b8 >> 8

			sum := r + g + b
			percent := int((sum * 100) / max)

			floored := (len(density) * percent) / 100

			char := density[floored]

			fmt.Printf("%s", string(char))
			f.Write([]byte(string(char)))
		}
		f.Write([]byte("\n"))
		fmt.Printf("\n")
	}
}
