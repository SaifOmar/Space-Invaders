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
	EnemyLevel
	health int
	color  color.RGBA
}

func (b *BaseEnemy) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.X), float32(b.Y), float32(b.Width), float32(b.Height), b.color, true)
	return
}

func (b *BaseEnemy) GetPosition() {
}

func (b *BaseEnemy) Update() {
}
