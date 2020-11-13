package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_3x3_space(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-mono-3x3")
	is.Nil(err)
	pix := draw(face, " ")
	exp := make([]uint8, (face.Width-2)*(face.Height-2))
	is.Equal(pix, exp, reprPix(pix, 3, 3))
}

func Test_3x3_a(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-mono-3x3")
	is.Nil(err)
	pix := draw(face, "a")
	exp := []uint8{
		0, X, X,
		X, 0, X,
		0, X, X,
	}
	is.Equal(pix, exp, reprPix(pix, 3, 3))
}

func Test_3x3_A(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-mono-3x3")
	is.Nil(err)
	pix := draw(face, "A")
	exp := []uint8{
		0, X, 0,
		X, X, X,
		X, 0, X,
	}
	is.Equal(pix, exp, reprPix(pix, 3, 3))
}

func Test_3x3_q(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-mono-3x3")
	is.Nil(err)
	pix := draw(face, "q")
	exp := []uint8{
		X, X, 0,
		X, X, 0,
		0, X, X,
	}
	is.Equal(pix, exp, reprPix(pix, 3, 3))
}
