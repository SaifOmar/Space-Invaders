package weapon

import (
	"image/color"

	"githbu.com/SaifOmar/space-invaders/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Projectile struct {
	Speed int
	Dmg   int
	Color color.RGBA
	*entity.Position
	*entity.Dimensions
}

func NewProjectile(dmg int, p entity.Position) *Projectile {
	return &Projectile{
		Speed:    3,
		Dmg:      dmg,
		Color:    color.RGBA{255, 0, 0, 255},
		Position: p.ConstructNewPosition(),
		Dimensions: &entity.Dimensions{
			Width:  10,
			Height: 10,
		},
	}
}

func (p *Projectile) Update() {
	p.Y += p.Speed
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, 10, 10, 10, p.Color, true)
}
