package main

import "github.com/hajimehoshi/ebiten"

// Duck game person
type Duck struct {
	x      int
	y      int
	sprite Sprite
}

// Draw the duck
func (duck *Duck) Draw(game *Game, screen *ebiten.Image) {
	duck.sprite.Draw(screen)
}

// Update the duck
func (duck *Duck) Update(game *Game, screen *ebiten.Image) {
	duck.sprite.x = duck.x
	duck.sprite.y = duck.y
	duck.sprite.Draw(screen)
}
