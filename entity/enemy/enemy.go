package enemy

import (
	"image/color"

	"githbu.com/SaifOmar/space-invaders/entity"
	"githbu.com/SaifOmar/space-invaders/weapon"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type EnemyLevel string

const (
	fr5a    EnemyLevel = "fr5a"
	deek               = "deek"
	n3ama              = "n3ama"
	dinasor            = "dinasor"
)

type Enemy interface {
	Update()
	Draw(*ebiten.Image)
	GetPosition()
}
type BaseEnemy struct {
	weapon.Weapon
	*entity.Position
	*entity.Dimensions
	health int
	EnemyLevel
	color color.RGBA
}

func (b *BaseEnemy) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.X), float32(b.Y), float32(b.Width), float32(b.Height), b.color, true)
	return
}

func (b *BaseEnemy) GetPosition() {
}

func (b *BaseEnemy) Update() {
}

func NewFr5aEnemy(x, y int) *BaseEnemy {
	b := &BaseEnemy{
		Position:   &entity.Position{X: x, Y: y},
		Dimensions: &entity.Dimensions{Width: 10, Height: 10},
		health:     100,
		EnemyLevel: fr5a,
		color:      color.RGBA{0x44, 0x44, 0x43, 0x30},
	}
	return b
}

func NewDeekEnemy(x, y int) *BaseEnemy {
	b := &BaseEnemy{
		Position:   &entity.Position{X: x, Y: y},
		Dimensions: &entity.Dimensions{Width: 10, Height: 10},
		health:     100,
	}
	return b
}

// func NewN3amaEnemy(x, y int) *BaseEnemy {
// 	b := &BaseEnemy{
// 		Position:   &Position{x: x, y: y},
// 		Dimensions: Dimensions{Width: 10, Height: 10},
// 		health:     100,
// 	}
// 	return b
// }
//
// func NewDinasorEnemy(x, y int) *BaseEnemy {
// 	b := &BaseEnemy{
// 		Position:   &Position{x: x, y: y},
// 		Dimensions: Dimensions{Width: 10, Height: 10},
// 		health:     100,
// 	}
// 	return b
// }
