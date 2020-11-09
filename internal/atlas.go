package internal

import (
	"fmt"

	"github.com/faiface/pixel/text"
)

func MemFont(name string) (*text.Atlas, error) {
	assets, err := NewAssets()
	if err != nil {
		return nil, fmt.Errorf("cannot read assets: %v", err)
	}

	config, err := NewConfig(assets, name)
	if err != nil {
		return nil, fmt.Errorf("cannot read config: %v", err)
	}

	stream, err := assets.OpenTTF(name)
	if err != nil {
		return nil, fmt.Errorf("cannot read TTF file: %v", err)
	}

	pix, err := ReadPix(stream, config)
	if err != nil {
		return nil, fmt.Errorf("cannot read sprite: %v", err)
	}
	face := NewFace(pix, config)

	atlas := text.NewAtlas(face, text.ASCII)
	return atlas, nil
}
