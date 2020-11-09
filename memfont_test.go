package memfont_test

import (
	"testing"

	"github.com/faiface/pixel/text"
	"github.com/matryer/is"
	"github.com/orsinium-labs/memfont"
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
	is := is.New(t)
	for name, f := range funcs() {
		t.Run(name, func(t *testing.T) {
			is := is.New(t)
			atlas, err := f()
			is.True(atlas != nil)
			is.NoErr(err)
		})
	}
}
