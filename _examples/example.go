package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/orsinium-labs/memfont"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	atlas, err := memfont.Prop5x6()
	if err != nil {
		panic(err)
	}
	txt := text.New(pixel.V(50, 500), atlas)
	txt.Color = color.RGBA{0xff, 0xff, 0xff, 0xff}

	typed := "Hello, world!\n"
	for !win.Closed() {
		if typed != "" {
			_, err = txt.WriteString(typed)
			if err != nil {
				panic(err)
			}

			win.Clear(color.RGBA{0x00, 0x00, 0x00, 0x00})
			mat := pixel.IM
			scale := 8.0
			mat = mat.Scaled(pixel.ZV, scale)
			center := win.Bounds().Center().Sub(txt.Bounds().Center().Scaled(scale))
			mat = mat.Moved(center)
			txt.Draw(win, mat)
		}

		typed = win.Typed()
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
