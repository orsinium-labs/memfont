package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_5x5_space(t *testing.T) {
	is := require.New(t)
	face, err := MemFont("mem-prop-5x5")
	is.Nil(err)
	pix := draw(face, " ")
	exp := make([]uint8, (face.Width-2)*(face.Height-2))
	is.Equal(pix, exp)
}

func Test_5x5_a(t *testing.T) {
	is := require.New(t)
	face, err := MemFont("mem-prop-5x5")
	is.Nil(err)
	pix := draw(face, "a")
	exp := []uint8{
		0, 0, 0, 0, 0,
		0, X, X, 0, 0,
		X, 0, X, 0, 0,
		0, X, X, 0, 0,
		0, 0, 0, 0, 0,
	}
	is.Equal(pix, exp)
}

func Test_5x5_A(t *testing.T) {
	is := require.New(t)
	face, err := MemFont("mem-prop-5x5")
	is.Nil(err)
	pix := draw(face, "A")
	exp := []uint8{
		0, X, 0, 0, 0,
		X, 0, X, 0, 0,
		X, X, X, 0, 0,
		X, 0, X, 0, 0,
		0, 0, 0, 0, 0,
	}
	is.Equal(pix, exp)
}
