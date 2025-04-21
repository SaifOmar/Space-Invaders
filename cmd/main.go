package main

import (
	"log"

	"githbu.com/SaifOmar/space-invaders/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Sapce Invaders")
	if err := game.Start(); err != nil {
		log.Fatal("Can't start error", err)
	}
}
