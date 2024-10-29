package nineslice

import (
	"image"
	"iter"

	"github.com/hajimehoshi/ebiten/v2"
)

func Rects(srcBounds, srcInner, dstBounds, dstInner image.Rectangle) iter.Seq2[image.Rectangle, image.Rectangle] {
	return func(yield func(src, dst image.Rectangle) bool) {
		yield(
			image.Rect(srcBounds.Min.X, srcBounds.Min.Y, srcInner.Min.X, srcInner.Min.Y),
			image.Rect(dstBounds.Min.X, dstBounds.Min.Y, dstInner.Min.X, dstInner.Min.Y),
		)
	}
}

func ScaledFragments(img *ebiten.Image, srcInner, dstBounds, dstInner image.Rectangle) iter.Seq2[*ebiten.Image, ebiten.GeoM] {
	return func(yield func(img *ebiten.Image, geom ebiten.GeoM) bool) {
		for src, dst := range Rects(img.Bounds(), srcInner, dstBounds, dstInner) {
			sx := float64(dst.Dx()) / float64(src.Dx())
			sy := float64(dst.Dy()) / float64(src.Dy())
			geom := ebiten.GeoM{}
			geom.Scale(sx, sy)
			geom.Translate(float64(dst.Min.X), float64(dst.Min.Y))
			yield(img.SubImage(src).(*ebiten.Image), geom)
		}
	}
}

func TiledFragments(img *ebiten.Image, srcInner, dstBounds, dstInner image.Rectangle) iter.Seq2[*ebiten.Image, ebiten.GeoM] {
	return func(yield func(img *ebiten.Image, geom ebiten.GeoM) bool) {
		for src, dst := range Rects(img.Bounds(), srcInner, dstBounds, dstInner) {
			sx := float64(dst.Dx()) / float64(src.Dx())
			sy := float64(dst.Dy()) / float64(src.Dy())
			geom := ebiten.GeoM{}
			geom.Scale(sx, sy)
			geom.Translate(float64(dst.Min.X), float64(dst.Min.Y))
			yield(img.SubImage(src).(*ebiten.Image), geom)
		}
	}
}
