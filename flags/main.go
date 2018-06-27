package main

import (
	"flag"
	"fmt"
	"image/color"
	"strconv"
)

type colorValue struct {
	color.Color
}

func (c *colorValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 16, 64)

	if err != nil {
		return fmt.Errorf("not a color: %v", err)
	}

	b := uint8(v % 256)
	g := uint8((v / 256) % 256)
	r := uint8((v / (256 * 256)) % 256)
	c.Color = color.RGBA{R: r, G: g, B: b, A: 255}
	return nil
}

func (c *colorValue) String() string {
	var r, g, b, a uint32
	if c != nil && c.Color != nil {
		r, g, b, a = c.RGBA()
		r, g, b, a = r/256, b/256, b/256, a/256
	}

	return fmt.Sprintf("rgba(%v, %v, %v, %v)", r, g, b, a)
}

func flagColor(name string, value color.Color, usage string) color.Color {
	v := &colorValue{value}
	flag.Var(v, name, usage)
	return v
}

func main() {
	fg := flagColor("fg", color.White, "foreground color")
	bg := flagColor("bg", color.Black, "background color")
	flag.Parse()

	draw(fg, bg)
}

func draw(fg, bg color.Color) {
	fmt.Printf("drawing with foreground %v and background %v\n", fg, bg)
}
