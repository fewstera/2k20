// +build linux
package main

import (
	"fmt"
	"image"
	"image/draw"
	"time"

	"github.com/fewstera/2k20/pkg/nyecanvas"
	rgbmatrix "github.com/mcuadros/go-rpi-rgb-led-matrix"
)

const pixelW = 32
const pixelH = 16
const windowScale = 37

func main() {
	// create a new Matrix instance with the DefaultConfig
	config := &rgbmatrix.DefaultConfig
	config.Rows = 16
	config.HardwareMapping = "adafruit-hat"
	config.Brightness = 55
	m, _ := rgbmatrix.NewRGBLedMatrix(config)
	c := rgbmatrix.NewCanvas(m)
	defer c.Close() // don't forgot close the Matrix, if not your leds will remain ono

	nc, err := nyecanvas.New(pixelW, pixelH)
	if err != nil {
		panic(fmt.Sprintf("Error init'ing NYE canvas: %s", err))
	}

	for {
		draw.Draw(c, c.Bounds(), nc.Tick(), image.ZP, draw.Src)

		// don't forget call Render to display the new led status
		c.Render()
		time.Sleep(time.Duration(10) * time.Millisecond)
	}
}
