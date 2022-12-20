package ext

import (
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"learn-go/extend/game/ebiten1/conf"
	"learn-go/extend/game/ebiten1/tips"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeOver
	ModeSuccess
)

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

type Game struct {
	mode              Mode
	input             *Input
	config            *conf.Config
	ship              *tips.Ship
	bullets           map[*tips.Bullet]struct{}
	aliens            map[*tips.Alien]struct{}
	bulletInterval    int
	maxBulletNum      int
	lastBulletAddTime int
	lastAlienAddTime  int
	lastAlien         *tips.Alien
	failCount         int
	successCount      int
}

func (g *Game) init() {
	g.CreateFonts()
}

func NewGame(config *conf.Config) *Game {
	g := &Game{
		input:             &Input{msg: "hello Input"},
		config:            config,
		ship:              tips.NewShip(config),
		bullets:           make(map[*tips.Bullet]struct{}),
		aliens:            make(map[*tips.Alien]struct{}),
		lastBulletAddTime: 0,
		bulletInterval:    config.BulletInterval,
		maxBulletNum:      config.MaxBulletNum,
	}
	g.init()
	return g
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		if g.input.IsKeyPressed(g) {
			g.mode = ModeGame
		}
	case ModeGame:
		g.input.Update(g)
		for bullet := range g.bullets {
			if !bullet.MoveY() {
				delete(g.bullets, bullet)
			}
		}
		g.addAlien()
		for alien := range g.aliens {
			alien.MoveX(g.aliens)
			if !alien.MoveY() {
				g.addFailNum()
				delete(g.aliens, alien)
			}
		}
		g.checkCollision()
		g.checkShipCollision()
	case ModeOver, ModeSuccess:
		if g.input.IsKeyPressed(g) {
			g.init()
			g.mode = ModeTitle
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.config.BgColor)

	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"ALIEN INVASION"}
		texts = []string{"", "", "", "", "", "", "", "PRESS SPACE KEY", "", "OR LEFT MOUSE"}
	case ModeGame:
		g.ship.Draw(screen)
		for bullet := range g.bullets {
			bullet.Draw(screen)
		}
		for alien := range g.aliens {
			alien.Draw(screen)
		}
	case ModeOver:
		texts = []string{"", "GAME OVER!"}
	case ModeSuccess:
		texts = []string{"", "YOU WIN!"}
	}

	for i, l := range titleTexts {
		x := (g.config.ScreenWidth - len(l)*g.config.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*g.config.TitleFontSize, color.White)
	}
	for i, l := range texts {
		x := (g.config.ScreenWidth - len(l)*g.config.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*g.config.FontSize, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, ScreenHeight int) {
	return g.config.LayoutWidth, g.config.LayoutHeight
}

func (g *Game) addBullet(bullet *tips.Bullet) {
	now := int(time.Now().UnixMilli())
	if g.lastBulletAddTime == 0 || now-g.lastBulletAddTime > g.bulletInterval {
		if g.maxBulletNum < len(g.bullets) {
			//fmt.Printf("超出子弹限制 %d 。。。 \n", g.maxBulletNum)
		} else {
			g.bullets[bullet] = struct{}{}
			g.lastBulletAddTime = now
		}
	} else {
		//fmt.Printf("发射子弹太快 last=%d now=%d 。。。 \n", g.lastBulletAddTime, now)
	}
}

func (g *Game) addAlien() {
	now := int(time.Now().UnixMilli())
	if g.lastAlienAddTime == 0 || now-g.lastAlienAddTime > g.config.AlienInterval {
		var alien *tips.Alien
		// alien之间不能碰撞 尝试5次创建
		for createNum := 5; createNum > 0; createNum++ {
			alien = tips.NewAlien(g.config)
			if g.lastAlien == nil || g.checkNewAlien(alien) {
				g.aliens[alien] = struct{}{}
				g.lastAlienAddTime = now
				g.lastAlien = alien
				break
			}
			createNum--
		}
	}
}

func (g *Game) checkNewAlien(newAline *tips.Alien) bool {
	for alien := range g.aliens {
		if tips.CheckReactCollision(newAline.GetRect(), alien.GetRect()) {
			return false
		}
	}
	return true
}

func (g *Game) addFailNum() {
	g.failCount++
	if g.failCount > g.config.FailCount {
		g.mode = ModeOver
	}
}

func (g *Game) checkCollision() {
	for alien := range g.aliens {
		for bullet := range g.bullets {
			if tips.CheckReactCollision(alien.GetRect(), bullet.GetRect()) {
				g.successCount++
				if g.successCount > g.config.SuccessCount {
					g.mode = ModeSuccess
				}
				delete(g.bullets, bullet)
				delete(g.aliens, alien)
			}
		}
	}
}

//checkShipCollision 检测飞船是否被撞
func (g *Game) checkShipCollision() {
	rectShip := g.ship.GetRect()
	for alien := range g.aliens {
		if tips.CheckReactCollision(alien.GetRect(), rectShip) {
			g.addFailNum()
			delete(g.aliens, alien)
		}
	}
}

func (g *Game) CreateFonts() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.config.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.config.FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.config.SmallFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
