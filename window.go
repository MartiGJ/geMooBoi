package main

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	initScreenWidth  = 160
	initScreenHeight = 144
	initScreenScale  = 2
)

var pixels = make([]byte, 4*initScreenHeight*initScreenWidth)

// update is called every frame (1/60 [s]).
func update(screen *ebiten.Image) error {

	// Write your game's logical update.
	updateState()
	if ebiten.IsDrawingSkipped() {
		// When the game is running slowly, the rendering result
		// will not be adopted.
		return nil
	}

	// Write your game's rendering.
	screen.ReplacePixels(pixels)

	return nil
}
