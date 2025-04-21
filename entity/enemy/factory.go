package enemy

import "githbu.com/SaifOmar/space-invaders/entity"

var enemyStartingPos = entity.Position{X: 300, Y: 300}

type EnemyFactory interface {
	CreateEnemy(x, y int) Enemy
}

func CreateEnemies(x int) []Enemy {
	var enemies []Enemy
	f := NewEnemyFacade(x)
	for range f.Rows {
		e := f.CreateEnemy()
		enemies = append(enemies, e)
	}
	return enemies
}

type EnemyFacade struct {
	Rows         int
	Row          int
	RowLength    int
	SpaceBetween int
	LastPosition entity.Position
	MaxBounds    entity.Position
	EnemyLevel
}

func NewEnemyFacade(Rows int) *EnemyFacade {
	return &EnemyFacade{
		Rows:         Rows,
		Row:          1,
		SpaceBetween: 100,
		RowLength:    10,
		LastPosition: enemyStartingPos,
		MaxBounds: entity.Position{
			X: 400,
			Y: 400,
		},
		EnemyLevel: fr5a,
	}
}

func (e *EnemyFacade) CreateEnemy() Enemy {
	var createdEnemy Enemy
	switch e.EnemyLevel {
	case fr5a:
		{
			createdEnemy = Fr5aFactory{}.CreateEnemy(e.LastPosition.X, e.LastPosition.Y)
		}
	}
	e.UpdateLastPos()

	return createdEnemy
}

func (e *EnemyFacade) UpdateLastPos() {
	e.LastPosition.X += e.SpaceBetween
}
