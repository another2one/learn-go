package main

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"learn-go/common/funcs"
	"learn-go/extend/game/ebiten1/conf"
	"learn-go/extend/game/ebiten1/ext"
	"log"
)

func main() {
	config := conf.LoadConfig(funcs.ProjectPath + "extend/game/ebiten1/resources/assets/config.json")
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle(config.Title)
	if err := ebiten.RunGame(ext.NewGame(config)); err != nil {
		log.Fatal(err)
	}
}
