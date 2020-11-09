package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/orsinium-labs/memfont/embedded"
	"github.com/rakyll/statik/fs"
)

type Assets struct {
	fs http.FileSystem
}

func NewAssets() (*Assets, error) {
	fs, err := fs.New()
	if err != nil {
		return nil, err
	}
	c := Assets{fs: fs}
	return &c, nil
}

func (c *Assets) OpenTTF(path string) (http.File, error) {
	return c.fs.Open("/" + path + ".ttf")
}

func (c *Assets) JSON(path string, target interface{}) error {
	stream, err := c.fs.Open("/" + path + ".json")
	if err != nil {
		return fmt.Errorf("cannot find JSON file: %v", err)
	}
	b, err := ioutil.ReadAll(stream)
	if err != nil {
		return fmt.Errorf("cannot read JSON file: %v", err)
	}
	err = json.Unmarshal(b, target)
	if err != nil {
		return fmt.Errorf("cannot parse JSON file: %v", err)
	}
	return nil
}
