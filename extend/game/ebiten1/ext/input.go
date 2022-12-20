package ext

import (
	"learn-go/extend/game/ebiten1/tips"

	"github.com/hajimehoshi/ebiten/v2"
)

type Input struct {
	msg string
}

func (input *Input) Update(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.ship.MoveX(-1 * float64(g.config.ShipSpeedFactor))
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.ship.MoveX(1 * float64(g.config.ShipSpeedFactor))
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.ship.MoveY(-1 * float64(g.config.ShipSpeedFactor))
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.ship.MoveY(1 * float64(g.config.ShipSpeedFactor))
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if len(g.bullets) < g.config.MaxBulletNum {
			g.addBullet(tips.NewBullet(g.config, g.ship))
		}
	}
}

func (input *Input) IsKeyPressed(g *Game) bool {
	return ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}
