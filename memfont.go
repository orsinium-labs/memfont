package memfont

import (
	"github.com/faiface/pixel/text"
	"github.com/orsinium-labs/memfont/internal"
)

func Prop56() (*text.Atlas, error) {
	return internal.MemFont("mem-prop-5x6")
}
