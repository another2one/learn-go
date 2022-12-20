package main

import (
	_ "embed"
	"learn-go/extend/game/ebiten1/conf"
	"learn-go/extend/game/ebiten1/ext"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	config := conf.LoadConfig("../conf/config.json")
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle(config.Title)
	if err := ebiten.RunGame(ext.NewGame(config)); err != nil {
		log.Fatal(err)
	}
}
