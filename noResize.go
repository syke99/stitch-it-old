package main

import (
	"image"

	"github.com/disintegration/gift"
	dmc "github.com/syke99/go-c2dmc"
)

func noResize(img image.Image) ([][]string, map[string]int) {

	g := gift.New(
		gift.Pixelate(10),
	)

	resImg := image.NewRGBA(g.Bounds(img.Bounds()))

	g.Draw(resImg, img)

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	var dmcs [][]string

	var colMap = make(map[string]int)

	for y := 0; y < h; y++ {
		var row []string
		var tempMap = make(map[string]int)
		for x := 0; x < w; x++ {
			col := resImg.At(x, y)

			cb := dmc.NewColorBank()

			r, g, b := cb.RgbA(col)

			d, _ := cb.RgbToDmc(r, g, b)

			if i, ok := tempMap[d]; ok {
				i++
				tempMap[d] = i
			} else {
				tempMap[d] = 0
			}

			row = append(row, d)
			x++
		}

		if y == (h - 1) {
			for col, num := range tempMap {
				colMap[col] = num
			}
		}

		dmcs = append(dmcs, row)
		y++
	}

	return dmcs, colMap
}
