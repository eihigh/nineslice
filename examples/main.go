package main

import (
	"github.com/eihigh/nineslice"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	var img *ebiten.Image
	var screen *ebiten.Image
	for img, geom := range nineslice.ScaledFragments(img, i, d, di) {
		op := &ebiten.DrawImageOptions{GeoM: geom}
		op.GeoM.Translate(100, 200)
		screen.DrawImage(img, op)
	}
}
