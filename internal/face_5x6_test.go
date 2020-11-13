package internal

import (
	"image"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const X = uint8(0xFF)

func Test_5x6_Bounds(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-5x6", 2)
	is.Nil(err)

	b, _, _ := face.GlyphBounds('a')
	is.Equal(int(b.Min.X), 0)
	is.Equal(int(b.Max.X), 7*64)
	is.Equal(int(b.Min.Y), -6*64)
	is.Equal(int(b.Max.Y), 2*64)

	dot := fixed.Point26_6{X: 0, Y: 0}
	dr, mask, maskp, _, _ := face.Glyph(dot, 'a')
	is.Equal(dr.Min.X, 0)
	is.Equal(maskp.X, 0)
	is.Equal(maskp.Y, 97*8)
	is.Equal(mask.Bounds().Min.X, 0)
	is.Equal(mask.Bounds().Max.X, 7)
	is.Equal(mask.Bounds().Min.Y, 0)
	is.Equal(mask.Bounds().Max.Y, 64*32)
}

func draw(is *require.Assertions, face *Face, text string) []uint8 {
	dst := image.NewGray(image.Rect(0, 0, face.Width-2, face.Height-2))
	d := font.Drawer{
		Dst:  dst,
		Src:  image.White,
		Face: face,
		Dot:  fixed.P(0, face.Height-2),
	}

	d.DrawString(text)
	return dst.Pix
}

func Test_5x6_space(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-5x6", 2)
	is.Nil(err)
	pix := draw(is, face, " ")
	exp := make([]uint8, (face.Width-2)*(face.Height-2))
	is.Equal(pix, exp)
}

func Test_5x6_a(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-5x6", 2)
	is.Nil(err)
	pix := draw(is, face, "a")
	exp := []uint8{
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, X, X, 0,
		0, X, 0, X, 0,
		0, X, 0, X, 0,
		0, 0, X, X, 0,
	}
	is.Equal(pix, exp)
}

func Test_5x6_0(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-5x6", 2)
	is.Nil(err)
	pix := draw(is, face, "0")
	exp := []uint8{
		0, 0, 0, 0, 0,
		0, 0, X, X, 0,
		0, X, 0, X, 0,
		0, X, 0, X, 0,
		0, X, 0, X, 0,
		0, X, X, 0, 0,
	}
	is.Equal(pix, exp)
}

func Test_5x6_z(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-5x6", 2)
	is.Nil(err)
	pix := draw(is, face, "z")
	exp := []uint8{
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, X, X, 0, 0,
		0, 0, X, 0, 0,
		0, 0, X, 0, 0,
		0, 0, X, X, 0,
	}
	is.Equal(pix, exp)
}
