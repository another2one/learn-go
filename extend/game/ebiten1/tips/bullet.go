package tips

import (
	"image"
	"learn-go/extend/game/ebiten1/conf"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	image       *ebiten.Image
	width       int
	height      int
	x           float64
	y           float64
	speedFactor float64
	limit       Limit
}

func NewBullet(cfg *conf.Config, ship *Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)

	return &Bullet{
		image:       img,
		width:       cfg.BulletWidth,
		height:      cfg.BulletHeight,
		x:           ship.x + float64(ship.Width-cfg.BulletWidth)/2,
		y:           float64(cfg.ScreenHeight - ship.Height - cfg.BulletHeight),
		speedFactor: cfg.BulletSpeedFactor,
		limit:       Limit{0, float64(cfg.LayoutWidth - ship.Width), 0, float64(cfg.LayoutHeight - ship.Height)},
	}
}

func (bulet *Bullet) GetRect() Rect {
	return Rect{float64(bulet.width), float64(bulet.height), bulet.x, bulet.y}
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.x, bullet.y)
	screen.DrawImage(bullet.image, op)
}

func (bullet *Bullet) MoveY() bool {
	bullet.y -= bullet.speedFactor
	return bullet.y < bullet.limit.YMax && bullet.y > 0
}
