package conf

import (
	"encoding/json"
	"image/color"
	"learn-go/extend/game/ebiten1/resources"
	"log"
)

type Config struct {
	ScreenWidth       int        `json:"screenWidth"`
	ScreenHeight      int        `json:"screenHeight"`
	LayoutWidth       int        `json:"layoutWidth"`
	LayoutHeight      int        `json:"layoutHeight"`
	Title             string     `json:"title"`
	BgColor           color.RGBA `json:"bgColor"`
	ShipSpeedFactor   float64    `json:"shipSpeedFactor"`
	BulletWidth       int        `json:"bulletWidth"`
	BulletHeight      int        `json:"bulletHeight"`
	BulletSpeedFactor float64    `json:"bulletSpeedFactor"`
	BulletColor       color.RGBA `json:"bulletColor"`
	MaxBulletNum      int        `json:"maxBulletNum"`
	BulletInterval    int        `json:"bulletInterval"`
	AlienSpeedFactor  float64    `json:"alienSpeedFactor"`
	AlienSpeedFactorX float64    `json:"alienSpeedFactorX"`
	AlienInterval     int        `json:"alienInterval"`
	TitleFontSize     int        `json:"titleFontSize"`
	FontSize          int        `json:"fontSize"`
	SmallFontSize     int        `json:"smallFontSize"`
	SuccessCount      int        `json:"successCount"`
	FailCount         int        `json:"failCount"`
}

func LoadConfig(path string) *Config {
	config := &Config{}
	data, err := resources.EmbedPath.ReadFile("assets/config.json")
	if err != nil {
		log.Fatalf("读取 EmbedPath 文件错误: %v \n", err)
	}
	if err := json.Unmarshal(data, config); err != nil {
		log.Fatalf("文件 ./config.json Unmarshal 错误 %s", err)
	}

	config.LayoutWidth = config.ScreenWidth
	config.LayoutHeight = config.ScreenHeight
	return config
}
