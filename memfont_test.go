package memfont_test

import (
	"fmt"
	"testing"

	"github.com/orsinium-labs/memfont"
	"github.com/stretchr/testify/require"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type F = func() (font.Face, error)

func funcs() map[string]F {
	return map[string]F{
		"m-3x3": memfont.Mono3x3,
		"m-4x4": memfont.Mono4x4,
		"p-3x5": memfont.Prop3x5,
		"p-4x4": memfont.Prop4x4,
		"p-5x5": memfont.Prop5x5,
		"p-5x6": memfont.Prop5x6,
	}

}

func TestNoErrors(t *testing.T) {
	for name, f := range funcs() {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)
			face, err := f()
			is.NotNil(face)
			is.Nil(err)
		})
	}
}

func TestSizeRange(t *testing.T) {
	for name, f := range funcs() {
		for _, glyph := range "123456789qwertyQWERTYabcABC" {
			name = fmt.Sprintf("%s/%c", name, glyph)
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				is := require.New(t)
				face, err := f()
				is.Nil(err)
				dr, _, _, _, ok := face.Glyph(fixed.Point26_6{}, glyph)
				is.True(ok)
				w := dr.Dx()
				h := dr.Dy()

				is.Greater(w, 2)
				is.Greater(h, 2)
				is.Less(w, 9)
				is.Less(h, 9)
			})
		}
	}
}
