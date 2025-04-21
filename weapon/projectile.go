package weapon

import (
	"image/color"

	"githbu.com/SaifOmar/space-invaders/entity"
)

type Projectile struct {
	speed int
	dmg   int
	color color.RGBA
	*entity.Position
	*entity.Dimensions
}
