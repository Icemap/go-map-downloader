package main

import (
	"flag"
)

// MapConfig map config struct
type MapConfig struct {
	// GoroutineNum concurrency num
	GoroutineNum int
	// SavePath map save root path
	SavePath string
	// Combine put same level map together
	Combine bool
	// QPS query file speed per second
	QPS int

	// Map

	// LeftLongitude left longitude
	LeftLongitude float64
	// RightLongitude right longitude
	RightLongitude float64
	// TopLatitude top latitude
	TopLatitude float64
	// BottomLatitude bottom latitude
	BottomLatitude float64
	// MinLevel map min level
	MinLevel int
	// MaxLevel map max level
	MaxLevel int
	// MapType only in GoogleSatellite / GoogleImage / GoogleTerrain / AMapSatellite / AMapCover / AMapImage
	MapType string
	// GoogleWithLabel only effect when the MapType is GoogleSatellite / GoogleImage / GoogleTerrain
	GoogleWithLabel bool

	// Retry

	// MaxRetryNum max retry num
	MaxRetryNum int
}

// parseMapConfig 解析配置
func parseMapConfig() MapConfig {
	conf := MapConfig{}

	// parse rule
	flag.StringVar(&conf.SavePath, "p", "/tmp", "map save path")
	flag.IntVar(&conf.GoroutineNum, "g", 50, "goroutine nums")
	flag.BoolVar(&conf.Combine, "c", true, "combine same level map together")
	flag.IntVar(&conf.QPS, "q", 500, "query file per second number")

	flag.Float64Var(&conf.LeftLongitude, "l", 0, "left longitude")
	flag.Float64Var(&conf.RightLongitude, "r", 0, "right longitude")
	flag.Float64Var(&conf.TopLatitude, "t", 0, "top latitude")
	flag.Float64Var(&conf.BottomLatitude, "b", 0, "bottom latitude")
	flag.IntVar(&conf.MinLevel, "min", 1, "map min level")
	flag.IntVar(&conf.MaxLevel, "max", 3, "map max level")

	flag.StringVar(&conf.MapType, "type", "GoogleSatellite", "map type (GoogleSatellite/GoogleImage/GoogleTerrain/AMapSatellite/AMapCover/AMapImage)")
	flag.BoolVar(&conf.GoogleWithLabel, "google-label", true, "only effect when the map type is GoogleSatellite / GoogleImage / GoogleTerrain")

	flag.IntVar(&conf.MaxRetryNum, "retry", 3, "max retry num")

	// parse
	flag.Parse()

	return conf
}
