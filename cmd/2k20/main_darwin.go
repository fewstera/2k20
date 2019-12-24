// +build darwin
package main

import (
	"fmt"
	"image"
	"time"

	"github.com/fewstera/2k20/pkg/nyecanvas"
	"github.com/fogleman/gg"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

const pixelW = 32
const pixelH = 16
const windowScale = 37

func main() {
	windowW := pixelW * windowScale
	windowH := pixelH * windowScale
	wnd, cv, err := sdlcanvas.CreateWindow(windowW, windowH, "Hello")
	if err != nil {
		panic(err)
	}
	defer wnd.Destroy()

	nc, err := nyecanvas.New(pixelW, pixelH)
	if err != nil {
		panic(fmt.Sprintf("Error init'ing NYE canvas: %s", err))
	}

	wnd.MainLoop(func() {
		tickImage := nc.Tick()

		sc := scalePixelImage(tickImage, int(windowScale))

		cv.SetFillStyle("#fff")
		cv.FillRect(0, 0, float64(windowW), float64(windowH))
		cv.DrawImage(sc, 0, 0)
		time.Sleep(time.Millisecond * time.Duration(30))
	})
}

func scalePixelImage(in image.Image, scale int) image.Image {
	inW := in.Bounds().Dx()
	inH := in.Bounds().Dy()
	scW := inW * scale
	scH := inH * scale
	sc := gg.NewContext(scW, scH)
	for x := 0; x < inW; x++ {
		for y := 0; y < inH; y++ {
			sc.SetColor(in.At(x, y))
			sc.DrawRectangle(float64(x*scale), float64(y*scale), float64(scale), float64(scale))
			sc.Fill()
		}
	}
	return sc.Image()
}
