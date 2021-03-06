package memfont

import (
	"github.com/orsinium-labs/memfont/internal"
	"golang.org/x/image/font"
)

func Mono3x3() (font.Face, error) {
	return internal.MemFont("mem-mono-3x3")
}

func Mono4x4() (font.Face, error) {
	return internal.MemFont("mem-mono-4x4")
}

func Prop3x5() (font.Face, error) {
	return internal.MemFont("mem-prop-3x5")
}

func Prop4x4() (font.Face, error) {
	return internal.MemFont("mem-prop-4x4")
}

func Prop5x5() (font.Face, error) {
	return internal.MemFont("mem-prop-5x5")
}

func Prop5x6() (font.Face, error) {
	return internal.MemFont("mem-prop-5x6")
}
