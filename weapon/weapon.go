package weapon

import (
	"githbu.com/SaifOmar/space-invaders/entity"
)

type Weapon interface {
	Fire(P entity.Position, dt float64) *Projectile
	Update(dt float64)
}

type BaseWeapon struct {
	Cooldown    float64
	Timeelapsed float64
}

func (w *BaseWeapon) Fire(p entity.Position, dt float64) *Projectile {
	return NewProjectile(10, p)
}

func (w *BaseWeapon) Update(dt float64) {
	w.Timeelapsed += dt
}
