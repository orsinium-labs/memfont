package memfont_test

import (
	"fmt"
	"testing"

	"github.com/faiface/pixel/text"
	"github.com/orsinium-labs/memfont"
	"github.com/stretchr/testify/require"
)

type F = func() (*text.Atlas, error)

func funcs() map[string]F {
	return map[string]F{
		"3x5": memfont.Prop3x5,
		"4x4": memfont.Prop4x4,
		"5x5": memfont.Prop5x5,
		"5x6": memfont.Prop5x6,
	}

}

func TestNoErrors(t *testing.T) {
	for name, f := range funcs() {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)
			atlas, err := f()
			is.NotNil(atlas)
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
				atlas, err := f()
				is.Nil(err)
				is.True(atlas.Contains(glyph))
				g := atlas.Glyph(glyph)
				w := g.Frame.W()
				h := g.Frame.H()

				is.Greater(w, 2.0)
				is.Greater(h, 2.0)
				is.Less(w, 9.0)
				is.Less(h, 9.0)
			})
		}
	}
}
