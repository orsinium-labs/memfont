# mem-font -> pixel

A simple [Go](https://golang.org/) library that allows you to use [mem-font](https://github.com/oddoid/mem) pixel font.

It returns [font.Face](https://godoc.org/golang.org/x/image/font#Face) interface. A few ways how to use it:

+ Draw text on an image using [font.Drawer](https://godoc.org/golang.org/x/image/font#Drawer).
+ Use font in [pixel](https://github.com/faiface/pixel) game library via [pixel.Atlas](https://pkg.go.dev/github.com/faiface/pixel/text#Atlas).

See [_examples](./_examples/) directory.

```bash
go get github.com/orsinium-labs/memfont
```

<img src="./_examples/image/result.png" style="min-width:50%" />
