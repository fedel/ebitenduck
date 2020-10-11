package main

import (
	"bytes"
	"image"
	"log"

	"github.com/fedel/ebitenduck/resources/images"
	"github.com/hajimehoshi/ebiten"
)

var (
	duckImage *ebiten.Image
)

// LoadResources from files
func loadResources() {
	var err error

	img, _, err := image.Decode(bytes.NewReader(images.Duck_png))
	if err != nil {
		log.Fatal(err)
	}
	duckImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

// Game implements ebiten.Game interface.
type Game struct {
	inited bool
	duck   *Duck
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	g.duck = &Duck{
		x:      0,
		y:      0,
		sprite: MakeSprite(duckImage),
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	if !g.inited {
		g.init()
	}

	// Write your game's logical update.
	g.duck.x++
	g.duck.Update(g, screen)

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.duck.Draw(g, screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
