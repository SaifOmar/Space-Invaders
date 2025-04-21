package entity

type Dimensions struct {
	Width  int
	Height int
}

type Position struct {
	X int
	Y int
}

func (p Position) ConstructNewPosition() *Position {
	return &Position{
		X: p.X,
		Y: p.Y + 20,
	}
}
