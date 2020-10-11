package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	loadResources()

	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Duck Go Duck")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
