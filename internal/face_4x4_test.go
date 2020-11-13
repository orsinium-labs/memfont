package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_4x4_space(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-4x4", 0)
	is.Nil(err)
	pix := draw(face, " ")
	exp := make([]uint8, (face.Width-2)*(face.Height-2))
	is.Equal(pix, exp)
}

func reprPix(pix []uint8) string {
	s := ""
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if pix[i*4+j] == 0 {
				s += "."
			} else {
				s += "X"
			}
		}
		s += "\n"
	}
	return s
}

func Test_4x4_a(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-4x4", 2)
	is.Nil(err)
	pix := draw(face, "a")
	exp := []uint8{
		0, 0, 0, 0,
		0, 0, X, X,
		0, X, 0, X,
		0, 0, X, X,
	}
	is.Equal(pix, exp, reprPix(pix))
}

func Test_4x4_A(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-4x4", 2)
	is.Nil(err)
	pix := draw(face, "A")
	exp := []uint8{
		0, 0, X, 0,
		0, X, 0, X,
		0, X, X, X,
		0, X, 0, X,
	}
	is.Equal(pix, exp, reprPix(pix))
}
