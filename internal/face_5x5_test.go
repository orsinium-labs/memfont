package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_5x5_space(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-5x5", 2)
	is.Nil(err)
	pix := draw(is, face, " ")
	exp := make([]uint8, (face.Width-2)*(face.Height-2))
	is.Equal(pix, exp)
}

func Test_5x5_a(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-5x5", 2)
	is.Nil(err)
	pix := draw(is, face, "a")
	exp := []uint8{
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, X, X, 0,
		0, X, 0, X, 0,
		0, 0, X, X, 0,
	}
	is.Equal(pix, exp)
}

func Test_5x5_A(t *testing.T) {
	is := require.New(t)
	face, err := makeFace("mem-prop-5x5", 2)
	is.Nil(err)
	pix := draw(is, face, "A")
	exp := []uint8{
		0, 0, 0, 0, 0,
		0, 0, X, 0, 0,
		0, X, 0, X, 0,
		0, X, X, X, 0,
		0, X, 0, X, 0,
	}
	is.Equal(pix, exp)
}
