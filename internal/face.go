package internal

import (
	"image"
	"strings"

	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type Face struct {
	basicfont.Face
	config Config
}

func (f *Face) AdvanceFor(r rune) int {
	s := string(r)
	a, ok := f.config.Advances[s]
	if !ok {
		a = f.config.Advance
		if strings.ToUpper(s) == s {
			a += f.config.Leading
		}
	}
	return a
}

func (f *Face) Glyph(dot fixed.Point26_6, r rune) (
	image.Rectangle, image.Image, image.Point, fixed.Int26_6, bool,
) {
	advance := f.AdvanceFor(r)
	dr, mask, maskp, _, ok := f.Face.Glyph(dot, r)
	return dr, mask, maskp, fixed.I(advance + 1), ok
}

func (f *Face) GlyphBounds(r rune) (bounds fixed.Rectangle26_6, advance fixed.Int26_6, ok bool) {
	a := f.AdvanceFor(r)
	return fixed.R(0, -f.Ascent, f.Width, +f.Descent), fixed.I(a + 1), true
}

func NewFace(pix []byte, c Config) *Face {
	mask := &image.Alpha{
		Stride: c.Width,
		Rect:   image.Rectangle{Max: image.Point{c.Width, 256 * c.Height}},
		Pix:    pix,
	}
	bface := basicfont.Face{
		Advance: c.Advance,
		Width:   c.Width,
		Height:  c.Height,
		Ascent:  c.Height - c.Baseline,
		Descent: c.Baseline,
		Left:    -1,
		Mask:    mask,
		Ranges: []basicfont.Range{
			{Low: rune(0), High: rune(c.Count), Offset: 0},
			{Low: '\ufffd', High: '\ufffe', Offset: 0}, // unknown char
		},
	}
	face := Face{
		Face:   bface,
		config: c,
	}
	return &face
}
