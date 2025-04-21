package enemy

import (
	"image/color"

	"githbu.com/SaifOmar/space-invaders/entity"
)

type Fr5aFactory struct{}

func (f Fr5aFactory) CreateEnemy(x, y int) Enemy {
	b := &BaseEnemy{
		Position:   &entity.Position{X: x, Y: y},
		Dimensions: &entity.Dimensions{Width: 10, Height: 10},
		health:     100,
		EnemyLevel: fr5a,
		color:      color.RGBA{0x44, 0x44, 0x43, 0x30},
	}
	return b
}
