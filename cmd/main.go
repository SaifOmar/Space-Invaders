package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Position struct {
	x int
	y int
}

type Dimensions struct {
	Width  int
	Height int
}
type Player struct {
	Name string
	*Dimensions
	*Position
	Speed int
	Dmg   int
}
type EnemyLevel int

const (
	fr5a    EnemyLevel = iota
	deek               = 1
	n3ama              = 2
	dinasor            = 3
)

type Weapon interface {
	drawProjectile()
}
type Enemy struct {
	EnemyLevel
	*Dimensions
	*Position
	Speed int
	Weapon
}
type Game struct {
	Player *Player
}

func (p *Player) move() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
		{
			p.x -= p.Speed
		}
	case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
		{
			p.x += p.Speed
		}
	case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
		{
			p.y += p.Speed
		}
	case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
		{
			p.y -= p.Speed
		}
	}
}

func NewPlayer(name string, speed, dmg int) *Player {
	p := &Player{
		Speed: speed,
		Dmg:   dmg,
		Name:  name,
		Position: &Position{
			x: 640,
			y: 600,
		},
		Dimensions: &Dimensions{
			Width:  30,
			Height: 50,
		},
	}
	return p
}

func (g *Game) Update() error {
	g.Player.move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Player != nil {
		g.Player.darw(screen)
	} else {
		panic("Cannot load player")
	}
	ebitenutil.DebugPrintAt(screen, "Hello, World!", 640, 360)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	p := NewPlayer("Saif", 10, 10)
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Sapce Invaders")
	if err := ebiten.RunGame(&Game{Player: p}); err != nil {
		log.Fatal(err)
	}
}

// we are now drawing the player
// time to make it so that I take player pos from the player
// now how can I make him move ?
// it should be the update that is called every second 60 times
// we check for a keypress and if a key is pressed we will move with player speed in that direction

func (p *Player) darw(screen *ebiten.Image) {
	color := color.RGBA{0x80, 0x80, 0x80, 0xff}
	vector.DrawFilledRect(screen, float32(p.x), float32(p.y), float32(p.Width), float32(p.Height), color, true)
	return
}
