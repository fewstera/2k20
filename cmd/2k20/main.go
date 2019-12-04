package main

import (
	"fmt"
	"image"
	"image/draw"
	"io/ioutil"
	"math/rand"
	"time"

	rgbmatrix "github.com/mcuadros/go-rpi-rgb-led-matrix"

	"github.com/fogleman/gg"
	"github.com/zachomedia/go-bdf"
	"golang.org/x/image/font"
)

func main() {
	message := "Hello! This is pretty cool, huh?"
	font, err := fontFace()
	if err != nil {
		panic(fmt.Sprintf("Error: %s", err))
	}

	pixelCount := len(message)*5 + 34

	// create a new Matrix instance with the DefaultConfig
	config := &rgbmatrix.DefaultConfig
	config.Rows = 16
	config.HardwareMapping = "adafruit-hat"
	m, _ := rgbmatrix.NewRGBLedMatrix(config)
	c := rgbmatrix.NewCanvas(m)
	defer c.Close() // don't forgot close the Matrix, if not your leds will remain ono

	for {
		dc := gg.NewContext(pixelCount, 16)
		dc.SetFontFace(font)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		dc.SetRGB(rand.Float64(), rand.Float64(), rand.Float64())
		dc.DrawString(message, 1, 10)
		dc.Fill()

		for i := -32; i < pixelCount; i++ {
			cc := gg.NewContext(32, 16)
			cc.DrawImage(dc.Image(), -i, 0)

			// using the standard draw.Draw function we copy a white image onto the Canvas
			draw.Draw(c, c.Bounds(), cc.Image(), image.ZP, draw.Src)

			// don't forget call Render to display the new led status
			c.Render()
			time.Sleep(time.Duration(25) * time.Millisecond)
		}
	}
}

func fontFace() (font.Face, error) {
	fontInput, err := ioutil.ReadFile("./fonts/5x7.bdf")
	if err != nil {
		return nil, fmt.Errorf("reading font file: %s", err)
	}

	font, err := bdf.Parse(fontInput)
	if err != nil {
		return nil, fmt.Errorf("parsing font: %s", err)
	}

	return font.NewFace(), nil
}
