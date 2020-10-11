package main

import "github.com/hajimehoshi/ebiten"

// Sprite structure
type Sprite struct {
	x      int
	y      int
	rotate float64
	scaleX float64
	scaleY float64
	op     ebiten.DrawImageOptions
	image  *ebiten.Image
}

// NewSprite Alloc a new Sprite in the heap
func NewSprite(image *ebiten.Image) *Sprite {
	return &Sprite{
		x:      0,
		y:      0,
		rotate: 0,
		scaleX: 1,
		scaleY: 1,
		op:     ebiten.DrawImageOptions{},
		image:  image,
	}
}

// MakeSprite Alloc a new Sprite in the heap
func MakeSprite(image *ebiten.Image) Sprite {
	return Sprite{
		x:      0,
		y:      0,
		rotate: 0,
		scaleX: 1,
		scaleY: 1,
		op:     ebiten.DrawImageOptions{},
		image:  image,
	}
}

// Draw the sprite on screen
func (s *Sprite) Draw(screen *ebiten.Image) {
	s.op.GeoM.Reset()
	s.op.GeoM.Scale(s.scaleX, s.scaleY)
	s.op.GeoM.Rotate(s.rotate)
	s.op.GeoM.Translate(float64(s.x), float64(s.y))

	screen.DrawImage(s.image, &s.op)
}
