package main

import (
	"fmt"
	"github.com/fogleman/gg"
)

const (
	// PerPicWidth per tile map pic width or height pixel num
	PerPicWidth = 256
)

func combine(xMin, xMax, yMin, yMax, z int, config MapConfig) error {
	fmt.Println(getCombinePicPath(config, z))

	resultMapPic := gg.NewContext((xMax - xMin + 1) * PerPicWidth, (yMax - yMin + 1) * PerPicWidth)

	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			itemImage, err := gg.LoadImage(getPath(config, x, y, z))
			if err != nil {
				fmt.Printf("can not load pic, error: %+v\n", err)
			}

			resultMapPic.DrawImage(itemImage, (x - xMin) * PerPicWidth, (y - yMin) * PerPicWidth)
		}
	}

	err := resultMapPic.SavePNG(getCombinePicPath(config, z))
	if err != nil {
		return err
	}
	return nil
}
