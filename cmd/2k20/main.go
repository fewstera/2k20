package main

import (
	"image"
	"image/color"
	"image/draw"
	"time"

	rgbmatrix "github.com/mcuadros/go-rpi-rgb-led-matrix"
)

func main() {
	// create a new Matrix instance with the DefaultConfig
	config := &rgbmatrix.DefaultConfig
	config.Rows = 16
	config.HardwareMapping = "adafruit-hat"
	m, _ := rgbmatrix.NewRGBLedMatrix(config)

	// create the Canvas, implements the image.Image interface
	c := rgbmatrix.NewCanvas(m)
	defer c.Close() // don't forgot close the Matrix, if not your leds will remain on

	// using the standard draw.Draw function we copy a white image onto the Canvas
	draw.Draw(c, c.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// don't forget call Render to display the new led status
	c.Render()
	time.Sleep(time.Duration(10) * time.Second)
}
