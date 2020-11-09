package memfont

import (
	"github.com/faiface/pixel/text"
	"github.com/orsinium-labs/memfont/internal"
)

func Prop3x5() (*text.Atlas, error) {
	return internal.MemFont("mem-prop-3x5")
}

func Prop4x4() (*text.Atlas, error) {
	return internal.MemFont("mem-prop-4x4")
}

func Prop5x5() (*text.Atlas, error) {
	return internal.MemFont("mem-prop-5x5")
}

func Prop5x6() (*text.Atlas, error) {
	return internal.MemFont("mem-prop-5x6")
}
