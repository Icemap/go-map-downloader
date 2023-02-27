# go-map-downloader

![build_badge](https://github.com/Icemap/go-map-downloader/workflows/Go/badge.svg)

[English](README.md) | 中文

Golang编写的地图下载器. 支持多种地图类型:

- 谷歌卫星图
- 谷歌标准地图
- 谷歌地形图
- 高德卫星图
- 高德覆盖层图
- 高德标准地图

## 功能

- 下载地图瓦片
- 拼接瓦片为大地图

## 安装

### Gitpod（推荐）

[![gitpod_badge](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/Icemap/go-map-downloader)

你可以点击上方，在 Gitpod 中打开这个项目。Gitpod 是一个完整的云开发环境，你可以把它当作是远程的 VSCode 来使用。这使得你无需配置本地环境（而且这个环境是翻了墙的）。

因为已经运行了编译命令，可以直接使用二进制文件：`bin/go-map-downloader` 。 尝试运行：

```bash
./bin/go-map-downloader -l 139.278433 -t 35.968355 -r 140.506452 -b 35.427143 -min 11 -max 11 -type GoogleSatellite -p bin/save
```

结果如下：

![gitpod_result](./pic/gitpod.png)

### 本地安装

```bash
go get -u github.com/Icemap/go-map-downloader
```

## 🌰例子

### 谷歌卫星图

```bash
./go-map-downloader -l 139.278433 -t 35.968355 -r 140.506452 -b 35.427143 -min 11 -max 11 -type GoogleSatellite
```

![google satellite](pic/google_satellite_level_11.jpg)

> **Note:**
>
> 你可以使用 `google-label` 参数控制是否隐藏 Google 类型地图的标签。如：
> 
> ```
> ./go-map-downloader -l 139.278433 -t 35.968355 -r 140.506452 -b 35.427143 -min 11 -max 11 -type GoogleSatellite -google-label=false
> ```
> 
> 请注意 `google-label` 标签仅在 Google 类型地图中有效。

### 高德标准地图

```bash
./go-map-downloader -l 139.278433 -t 35.968355 -r 140.506452 -b 35.427143 -min 11 -max 11 -type AMapImage
```

![amap_image](pic/amap_image_level_11.jpg)

### 帮助

```
./go-map-downloader -h
Usage of ./go-map-downloader:
  -b float
        bottom latitude
  -c    combine same level map together (default true)
  -g int
        goroutine nums (default 50)
  -google-label
        only effect when the map type is GoogleSatellite / GoogleImage / GoogleTerrain (default true)
  -l float
        left longitude
  -max int
        map max level (default 3)
  -min int
        map min level (default 1)
  -p string
        map save path (default "/tmp")
  -q int
        query file per second number (default 500)
  -r float
        right longitude
  -retry int
        max retry num (default 3)
  -t float
        top latitude
  -type string
        map type (GoogleSatellite/GoogleImage/GoogleTerrain/AMapSatellite/AMapCover/AMapImage) (default "GoogleSatellite")
```