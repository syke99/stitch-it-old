package main

import (
	"image"
	"strconv"
	"strings"

	"github.com/disintegration/gift"
	dmc "github.com/syke99/go-c2dmc"
)

func resizeImage(img image.Image, size string) [][]string {

	sizes := strings.Split(size, "x")

	iw, _ := strconv.Atoi(sizes[0])
	ih, _ := strconv.Atoi(sizes[1])

	g := gift.New(
		gift.Resize(iw, ih, gift.LanczosResampling),
	)

	res := image.NewRGBA(g.Bounds(img.Bounds()))

	g.Draw(res, img)

	gi := gift.New(
		gift.Pixelate(10),
	)

	resImg := image.NewRGBA(gi.Bounds(res.Bounds()))

	gi.Draw(resImg, res)

	bounds := resImg.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	var dmcs [][]string

	for y := 0; y < h; y++ {
		var row []string
		for x := 0; x < w; x++ {
			col := resImg.At(x, y)

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
