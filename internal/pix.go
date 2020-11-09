package internal

import (
	"image"
	"io"
)

func ReadPix(stream io.Reader, c Config) ([]byte, error) {
	img, _, err := image.Decode(stream)
	if err != nil {
		return nil, err
	}

	// imgH := img.Bounds().Max.Y
	result := make([]byte, 0)
	for glyph := 0; glyph < c.Count; glyph++ {
		cornerX := (glyph % 16) * c.Width
		cornerY := glyph / 16 * c.Height
		for y := cornerY; y < cornerY+c.Height; y++ {
			for x := cornerX; x < cornerX+c.Width; x++ {
				r, g, b, a := img.At(x, y).RGBA()
				if r+g+b+a > 1000 {
					result = append(result, 0xff)
				} else {
					result = append(result, 0x00)
				}
			}
		}
	}
	return result, nil
}
