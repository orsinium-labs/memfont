package internal

import (
	"errors"
	"fmt"
)

func MemFont(name string) (*Face, error) {
	assets, err := NewAssets()
	if err != nil {
		return nil, fmt.Errorf("cannot read assets: %v", err)
	}

	config, err := NewConfig(assets, name)
	if err != nil {
		return nil, fmt.Errorf("cannot read config: %v", err)
	}

	stream, err := assets.OpenPNG(name)
	if err != nil {
		return nil, fmt.Errorf("cannot read PNG file: %v", err)
	}

	pix, err := ReadPix(stream, config)
	if err != nil {
		return nil, fmt.Errorf("cannot read sprite: %v", err)
	}
	if len(pix) < 600 {
		return nil, errors.New("not enough pixels")
	}
	face := NewFace(pix, config)
	return face, nil
}
