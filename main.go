package main

import (
	"fmt"
	"github.com/Icemap/coordinate"
	"github.com/cheggaaa/pb/v3"
	"github.com/panjf2000/ants/v2"
	"math"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Task struct {
	err      error
	x        int
	y        int
	z        int
	retryNum int
	config   MapConfig
}

func main() {
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)

	config := parseMapConfig()
	fmt.Printf("config: %+v\n", config)

	fmt.Println("start calculate...")
	// calculate
	leftTop, err := coordinate.Convert(coordinate.WGS84, coordinate.WebMercator, coordinate.Coordinate{X: config.LeftLongitude, Y: config.TopLatitude})
	if err != nil {
		panic(fmt.Errorf("left top error: %+v", err))
	}

	rightBottom, err := coordinate.Convert(coordinate.WGS84, coordinate.WebMercator, coordinate.Coordinate{X: config.RightLongitude, Y: config.BottomLatitude})
	if err != nil {
		panic(fmt.Errorf("right bottom error: %+v", err))
	}

	fmt.Printf("leftTop: %+v, rightBottom: %+v\n", leftTop, rightBottom)

	totalTask := 0
	for z := config.MinLevel; z <= config.MaxLevel; z++ {
		perTileWidth := coordinate.WebMercatorWidth / math.Pow(2, float64(z))
		leftTileX := int(leftTop.X / perTileWidth)
		rightTileX := int(rightBottom.X / perTileWidth)
		topTileY := int(leftTop.Y / perTileWidth)
		bottomTileY := int(rightBottom.Y / perTileWidth)
		totalTask += (rightTileX - leftTileX + 1) * (bottomTileY - topTileY + 1)
	}

	// download
	fmt.Println("start download...")
	bar := pb.StartNew(totalTask)
	var wg sync.WaitGroup

	pool, _ := ants.NewPoolWithFunc(config.GoroutineNum, func(iTask interface{}) {
		defer func() {
			wg.Done()
			bar.Increment()
		}()

		task, _ := iTask.(Task)
		err := getPic(task)
		for i := 0; i < task.config.MaxRetryNum; i++ {
			if err != nil {
				fmt.Printf("retry num: %d, task: %+v\n", i, task)
				err = getPic(task)
			}
		}
	})
	defer pool.Release()

	currentTask := 0
	for z := config.MinLevel; z <= config.MaxLevel; z++ {
		perTileWidth := coordinate.WebMercatorWidth / math.Pow(2, float64(z))
		leftTileX := int(leftTop.X / perTileWidth)
		rightTileX := int(rightBottom.X / perTileWidth)
		topTileY := int(leftTop.Y / perTileWidth)
		bottomTileY := int(rightBottom.Y / perTileWidth)

		currentTask++
		if currentTask%config.QPS == 0 {
			time.Sleep(time.Second)
		}

		for x := leftTileX; x <= rightTileX; x++ {
			for y := topTileY; y <= bottomTileY; y++ {
				wg.Add(1)
				pool.Invoke(Task{
					x:        x,
					y:        y,
					z:        z,
					retryNum: 0,
					config:   config,
				})
			}
		}
	}

	wg.Wait()
	bar.Finish()

	// combine
	fmt.Println("start combine...")
	if !config.Combine {
		return
	}
	for z := config.MinLevel; z <= config.MaxLevel; z++ {
		perTileWidth := coordinate.WebMercatorWidth / math.Pow(2, float64(z))
		leftTileX := int(leftTop.X / perTileWidth)
		rightTileX := int(rightBottom.X / perTileWidth)
		topTileY := int(leftTop.Y / perTileWidth)
		bottomTileY := int(rightBottom.Y / perTileWidth)

		err := combine(leftTileX, rightTileX, topTileY, bottomTileY, z, config)
		if err != nil {
			fmt.Printf("combine pic error: %+v\n", err)
		}
	}
}

func getPic(task Task) error {
	if task.retryNum >= task.config.MaxRetryNum {
		fmt.Printf("task retry reached retry max time, err: %+v, task info: %+v\n", task.err, task)
		return nil
	}

	filePath := getPath(task.config, task.x, task.y, task.z)
	url := coordinate.WebMercatorTileToURLWithTiltStyle(task.config.MapType, task.x, task.y, task.z,
		coordinate.TiltStyle{GoogleWithLabel: task.config.GoogleWithLabel})

	return Download(url, filePath)
}

func getPath(config MapConfig, x, y, z int) string {
	return strings.Join([]string{
		config.SavePath, config.MapType, strconv.Itoa(z), strconv.Itoa(x), strconv.Itoa(y),
	}, string(filepath.Separator)) + ".jpg"
}

func getCombinePicPath(config MapConfig, z int) string {
	return strings.Join([]string{
		config.SavePath, config.MapType, "level_" + strconv.Itoa(z),
	}, string(filepath.Separator)) + ".jpg"
}
