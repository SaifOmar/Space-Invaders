package game

import (
	"githbu.com/SaifOmar/space-invaders/entity"
	"githbu.com/SaifOmar/space-invaders/entity/enemy"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player  *entity.Player
	Enemies []enemy.Enemy
}

func (g *Game) Update() error {
	g.Player.Move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Player != nil {
		g.Player.Darw(screen)
		if g.Enemies != nil {
			for _, enemy := range g.Enemies {
				enemy.Draw(screen)
			}
		} else {
			panic("Cannot load player")
		}
	} else {
		panic("Cannot load player")
	}
	ebitenutil.DebugPrintAt(screen, "Hello, World!", 640, 360)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func Start() error {
	p := entity.NewPlayer("Saif", 10, 100)
	var enemies []enemy.Enemy
	for range 4 {
		e := enemy.NewFr5aEnemy(100, 100)
		enemies = append(enemies, e)
	}
	g := &Game{p, enemies}
	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	return nil
}
