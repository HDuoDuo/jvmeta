package imageutil

import (
	"image"
	"image/draw"
)

func Watermark(src image.Image, wmk image.Image, pt image.Point) image.Image {
	dst := image.NewNRGBA(image.Rect(0, 0, src.Bounds().Dx(), src.Bounds().Dy()))
	draw.Draw(dst, dst.Bounds(), src, src.Bounds().Min, draw.Src)
	mark := Resize(wmk, 0, src.Bounds().Dy()/9)
	draw.Draw(dst, mark.Bounds().Add(dst.Bounds().Max.Sub(mark.Bounds().Max.Sub(image.Point{12, 0}))), mark, mark.Bounds().Min.Add(pt), draw.Over)
	return dst
}
