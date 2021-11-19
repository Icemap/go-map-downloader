package main

import (
	"math"
)

const (
	// WebMercatorWidth web mercator projection width or height
	WebMercatorWidth = WebMercatorHalfWidth * 2

	// WebMercatorHalfWidth half of web mercator projection width or height
	WebMercatorHalfWidth = 20037508.3427892

	// RoundDegree a round degree num
	RoundDegree = 360.0

	// HalfRoundDegree half round degree num
	HalfRoundDegree = RoundDegree / 2

	// QuarterRoundDegree quarter round degree num
	QuarterRoundDegree = HalfRoundDegree / 2

	// DegreeToRadian degree to radian constant
	DegreeToRadian = math.Pi / 180.0
)

type (
	// Coordinate point struct
	Coordinate struct {
		X float64
		Y float64
	}
)

func LonLatToGoogleMercator(lon, lat float64) Coordinate {
	toX := lon * WebMercatorHalfWidth / HalfRoundDegree + WebMercatorHalfWidth
	toY := math.Log(math.Tan((QuarterRoundDegree + lat) * math.Pi / RoundDegree)) / DegreeToRadian
	toY = toY * WebMercatorHalfWidth / HalfRoundDegree
	toY = WebMercatorHalfWidth - toY
	return Coordinate{X: toX, Y: toY}
}