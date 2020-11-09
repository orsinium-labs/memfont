package internal

type Config struct {
	Count    int
	Width    int            `json:"cel_width"`
	Height   int            `json:"cel_height"`
	Leading  int            `json:"leading"`
	Kerning  int            `json:"default_kerning"`
	Advances map[string]int `json:"letter_width"`
	Advance  int            `json:"default_letter_width"`
}

func NewConfig(assets *Assets, path string) (Config, error) {
	c := Config{Count: 128}
	err := assets.JSON(path, &c)
	c.Width += 2
	c.Height += 2
	return c, err
}
