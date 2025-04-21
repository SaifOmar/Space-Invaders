package game

import (
	"time"

	"githbu.com/SaifOmar/space-invaders/entity"
	"githbu.com/SaifOmar/space-invaders/entity/enemy"
	"githbu.com/SaifOmar/space-invaders/weapon"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player  *entity.Player
	Enemies []enemy.Enemy
	Projs   []*weapon.Projectile
}

func (g *Game) Update() error {
	g.Player.Move()
	g.UpdateEnemies()
	return nil
}

func (g *Game) UpdateEnemies() {
	now := time.Now()
	for _, e := range g.Enemies {
		e.Update()
		if proj := e.Fire(float64(now.Weekday())); proj != nil {
			g.Projs = append(g.Projs, proj)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Player != nil {
		g.Player.Darw(screen)
		if g.Enemies != nil {
			for _, enemy := range g.Enemies {
				enemy.Draw(screen)
			}
			if g.Projs != nil {
				for _, proj := range g.Projs {
					proj.Draw(screen)
				}
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
	enemies := enemy.CreateEnemies(3)
	g := &Game{p, enemies, nil}
	if err := ebiten.RunGame(g); err != nil {
		return err
	}
	return nil
}
