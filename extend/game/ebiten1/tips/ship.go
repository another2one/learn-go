package tips

import (
	"bytes"
	_ "image/png"
	"learn-go/extend/game/ebiten1/conf"
	"learn-go/extend/game/ebiten1/resources"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ship struct {
	Image         *ebiten.Image
	Width, Height int
	x, y          float64
	limit         Limit
}

func NewShip(config *conf.Config) *Ship {
	//img, _, err := ebitenutil.NewImageFromFile("../assets/feiji.png")
	data, err := resources.EmbedPath.ReadFile("assets/feiji.png")
	if err != nil {
		log.Fatalf("读取 EmbedPath 文件错误: %v \n", err)
	}
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	// fmt.Println(width, " -- ", height)
	limit := Limit{0, float64(config.LayoutWidth - width), 0, float64(config.LayoutHeight - height)}
	ship := &Ship{
		Image:  img,
		Width:  width,
		Height: height,
		x:      limit.XMax / 2,
		y:      limit.YMax,
		limit:  limit,
	}

	return ship
}

func (ship *Ship) GetRect() Rect {
	return Rect{float64(ship.Width), float64(ship.Height), ship.x, ship.y}
}

func (ship *Ship) MoveX(x float64) {
	if ship.x+x > ship.limit.XMin && ship.x+x < ship.limit.XMax {
		ship.x += x
	}
}

func (ship *Ship) MoveY(y float64) {
	if ship.y+y > ship.limit.YMin && ship.y+y < ship.limit.YMax {
		ship.y += y
	}
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ship.x, ship.y)
	screen.DrawImage(ship.Image, op)
}
