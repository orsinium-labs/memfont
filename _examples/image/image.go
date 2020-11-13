package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/orsinium-labs/memfont"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	width  = 100
	height = 8
)

func main() {
	face, err := memfont.Prop5x6()
	if err != nil {
		log.Fatalf("NewFace: %v", err)
	}

	img := image.NewGray(image.Rect(0, 0, width, height*2+1))
	d := font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: face,
		Dot:  fixed.P(1, height),
	}
	d.DrawString("The quick brown fox")
	d.Dot = fixed.P(1, height*2)
	d.DrawString("jumps over the lazy dog")

	stream, err := os.Create("image.png")
	if err != nil {
		log.Fatalf("os.Create: %v", err)
	}
	err = png.Encode(stream, img)
	if err != nil {
		log.Fatalf("png.Encode: %v", err)
	}
}
