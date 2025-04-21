package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	Name string
	*Dimensions
	*Position
	Speed  int
	Dmg    int
	health int
}

func NewPlayer(name string, speed, dmg int) *Player {
	p := &Player{
		Speed: speed,
		Dmg:   dmg,
		Name:  name,
		Position: &Position{
			X: 640,
			Y: 600,
		},
		Dimensions: &Dimensions{
			Width:  30,
			Height: 50,
		},
		health: 100,
	}
	return p
}

func (p *Player) Darw(screen *ebiten.Image) {
	color := color.RGBA{0x80, 0x80, 0x80, 0xff}
	vector.DrawFilledRect(screen, float32(p.X), float32(p.Y), float32(p.Width), float32(p.Height), color, true)
	return
}

func (p *Player) Move() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
		{
			p.X -= p.Speed
		}
	case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
		{
			p.X += p.Speed
		}
	case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
		{
			p.Y += p.Speed
		}
	case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
		{
			p.Y -= p.Speed
		}
	}
}
