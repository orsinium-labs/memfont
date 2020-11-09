package memfont_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/memfont"
)

func TestNoErrors(t *testing.T) {
	is := is.New(t)

	atlas, err := memfont.Prop5x6()
	is.True(atlas != nil)
	is.NoErr(err)

	atlas, err = memfont.Prop5x5()
	is.True(atlas != nil)
	is.NoErr(err)

	atlas, err = memfont.Prop4x4()
	is.True(atlas != nil)
	is.NoErr(err)

	// atlas, err = memfont.Prop3x5()
	// is.True(atlas != nil)
	// is.NoErr(err)

}
