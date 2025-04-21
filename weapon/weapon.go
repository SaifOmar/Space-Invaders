package weapon

import (
	"image/color"

	"githbu.com/SaifOmar/space-invaders/entity"
)

type Weapon interface {
	fire(P entity.Position) *Projectile
	update(dt float32)
}

type BaseWeapon struct {
	cooldown    int
	timeelapsed int
}

func NewFr5aWeapon(x, y int) BaseWeapon {
	b := BaseWeapon{
		cooldown:    10,
		timeelapsed: 10,
	}
	return b
}

func (w *BaseWeapon) fire(p entity.Position) *Projectile {
	pr := &Projectile{
		speed:    3,
		dmg:      34,
		color:    color.RGBA{0x10, 0x10, 0x10, 0x10},
		Position: p.ConstructNewPosition(),
		Dimensions: &entity.Dimensions{
			Width:  10,
			Height: 10,
		},
	}

	return pr
}

func (w *BaseWeapon) update() {
}
