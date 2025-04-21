package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Player  *Player
	Enemies []*BaseEnemy
}

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
	Speed  int
	Dmg    int
	health int
}
type EnemyLevel string

const (
	fr5a    EnemyLevel = "fr5a"
	deek               = "deek"
	n3ama              = "n3ama"
	dinasor            = "dinasor"
)

type Weapon interface {
	fire(P Position) *Projectile
	update(dt float32)
}

type Enemy interface {
	update()
	draw(*ebiten.Image)
	position()
}
type BaseEnemy struct {
	Weapon
	*Position
	Dimensions
	health int
	EnemyLevel
	color color.RGBA
}

type Projectile struct {
	speed int
	dmg   int
	color color.RGBA
	*Position
	Dimensions
}

type BaseWeapon struct {
	cooldown    int
	timeelapsed int
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
		health: 100,
	}
	return p
}

func (p *Player) darw(screen *ebiten.Image) {
	color := color.RGBA{0x80, 0x80, 0x80, 0xff}
	vector.DrawFilledRect(screen, float32(p.x), float32(p.y), float32(p.Width), float32(p.Height), color, true)
	return
}

func NewFr5aWeapon(x, y int) BaseWeapon {
	b := BaseWeapon{
		cooldown:    10,
		timeelapsed: 10,
	}
	return b
}

func (w *BaseWeapon) fire(p Position) *Projectile {
	pr := &Projectile{
		speed:    3,
		dmg:      34,
		color:    color.RGBA{0x10, 0x10, 0x10, 0x10},
		Position: ConstructPosition(p),
		Dimensions: Dimensions{
			Width:  10,
			Height: 10,
		},
	}

	return pr
}

func ConstructPosition(p Position) *Position {
	return &Position{
		x: p.x,
		y: p.y + 20,
	}
}

func (w *BaseWeapon) update() {
}

func NewFr5aEnemy(x, y int) *BaseEnemy {
	b := &BaseEnemy{
		Position:   &Position{x: x, y: y},
		Dimensions: Dimensions{Width: 10, Height: 10},
		health:     100,
		EnemyLevel: fr5a,
		color:      color.RGBA{0x44, 0x44, 0x43, 0x30},
	}
	return b
}

func NewDeekEnemy(x, y int) *BaseEnemy {
	b := &BaseEnemy{
		Position:   &Position{x: x, y: y},
		Dimensions: Dimensions{Width: 10, Height: 10},
		health:     100,
	}
	return b
}

func NewN3amaEnemy(x, y int) *BaseEnemy {
	b := &BaseEnemy{
		Position:   &Position{x: x, y: y},
		Dimensions: Dimensions{Width: 10, Height: 10},
		health:     100,
	}
	return b
}

func NewDinasorEnemy(x, y int) *BaseEnemy {
	b := &BaseEnemy{
		Position:   &Position{x: x, y: y},
		Dimensions: Dimensions{Width: 10, Height: 10},
		health:     100,
	}
	return b
}

func (b *BaseEnemy) draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(b.Width), float32(b.Height), b.color, true)
	return
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

func (g *Game) Update() error {
	g.Player.move()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.Player != nil {
		g.Player.darw(screen)
		if g.Enemies != nil {
			for _, enemy := range g.Enemies {
				enemy.draw(screen)
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

// we are now drawing the player
// time to make it so that I take player pos from the player
// now how can I make him move ?
// it should be the update that is called every second 60 times
// we check for a keypress and if a key is pressed we will move with player speed in that direction

func main() {
	p := NewPlayer("Saif", 10, 100)
	var enemies []*BaseEnemy
	for range 4 {
		e := NewFr5aEnemy(100, 100)
		enemies = append(enemies, e)
	}
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Sapce Invaders")
	if err := ebiten.RunGame(&Game{Player: p, Enemies: enemies}); err != nil {
		log.Fatal(err)
	}
}
