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

func clip(xMinRatio, xMaxRatio, yMinRatio, yMaxRatio float64, z int, config MapConfig) error {
	fmt.Println(getClippedPicPath(config, z))

	im, err := gg.LoadImage(getCombinePicPath(config, z))
	if err != nil {
		return err
	}

	size := im.Bounds().Size()
	imageWidth := float64(size.X)
	imageHeight := float64(size.Y)
	canvasWidth := int((xMaxRatio - xMinRatio) * imageWidth)
	canvasHeight := int((yMaxRatio - yMinRatio) * imageHeight)

	overview := gg.NewContext(canvasWidth, canvasHeight)
	overview.DrawImageAnchored(im, 0, 0, xMinRatio, yMinRatio)

	err = overview.SavePNG(getClippedPicPath(config, z))
	if err != nil {
		return err
	}
	return nil
}
