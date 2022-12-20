package tips

import (
	"bytes"
	_ "image/png"
	"learn-go/extend/game/ebiten1/conf"
	"learn-go/extend/game/ebiten1/resources"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Alien struct {
	image         *ebiten.Image
	width, height int
	x, y          float64
	speedFactor   float64
	speedFactorX  float64
	limit         Limit
}

func NewAlien(cfg *conf.Config) *Alien {
	//img, _, err := ebitenutil.NewImageFromFile("../assets/alien.png")
	data, err := resources.EmbedPath.ReadFile("assets/alien.png")
	if err != nil {
		log.Fatalf("读取 EmbedPath 文件错误: %v \n", err)
	}
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	// fmt.Println(width, " -- ", height)
	limit := Limit{0, float64(cfg.LayoutWidth - width), 0, float64(cfg.LayoutHeight - height)}
	Alien := &Alien{
		image:        img,
		width:        width,
		height:       height,
		x:            rand.Float64() * limit.XMax,
		y:            0,
		speedFactor:  cfg.AlienSpeedFactor,
		speedFactorX: cfg.AlienSpeedFactorX,
		limit:        limit,
	}

	return Alien
}

func (alien *Alien) GetRect() Rect {
	return Rect{float64(alien.width), float64(alien.height), alien.x, alien.y}
}

func (Alien *Alien) MoveX(aliens map[*Alien]struct{}) {
	if Alien.x+Alien.speedFactorX > Alien.limit.XMin && Alien.x+Alien.speedFactorX < Alien.limit.XMax {
		Alien.x += Alien.speedFactorX
		Alien.speedFactorX = -Alien.speedFactorX
	}
}

func (alien *Alien) MoveY() bool {
	alien.y += alien.speedFactor
	return alien.y < alien.limit.YMax && alien.y > alien.limit.YMin
}

func (Alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(Alien.x, Alien.y)
	screen.DrawImage(Alien.image, op)
}
