package main

import (
	"image"

	"github.com/disintegration/gift"
	dmc "github.com/syke99/go-c2dmc"
)

func noResize(img image.Image) [][]string {

	gi := gift.New(
		gift.Pixelate(10),
	)

	resImg := image.NewRGBA(gi.Bounds(img.Bounds()))

	gi.Draw(resImg, img)

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	var dmcs [][]string

	for y := 0; y < h; y++ {
		var row []string
		for x := 0; x < w; x++ {
			col := img.At(x, y)

			cb := dmc.NewColorBank()

			r, g, b := cb.RgbA(col)

			d, _ := cb.RgbToDmc(r, g, b)

			row = append(row, d)
			x++
		}
		dmcs = append(dmcs, row)
		y++
	}

	return dmcs
}
