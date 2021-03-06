# memfont

A simple [Go](https://golang.org/) library providing interface for [mem-font](https://github.com/oddoid/mem) pixel font.

```bash
go get github.com/orsinium-labs/memfont
```

## Usage

It returns [font.Face](https://godoc.org/golang.org/x/image/font#Face) interface. A few ways how to use it:

+ Draw text on an image using [font.Drawer](https://godoc.org/golang.org/x/image/font#Drawer).
+ Use font in [pixel](https://github.com/faiface/pixel) game library via [pixel.Atlas](https://pkg.go.dev/github.com/faiface/pixel/text#Atlas).

See [_examples](./_examples/) directory.
