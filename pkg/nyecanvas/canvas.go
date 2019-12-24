package nyecanvas

import (
	"fmt"
	"image"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

type Screen int

const (
	CurrentTime Screen = iota
	MinSecRemaining
)

type Canvas struct {
	width     int
	height    int
	smallFont font.Face
	bigFont   font.Face
}

func New(width, height int) (*Canvas, error) {
	smallFont, err := loadFont("5x7")
	if err != nil {
		return nil, fmt.Errorf("loading small font: %s", err)
	}

	bigFont, err := loadFont("6x12")
	if err != nil {
		return nil, fmt.Errorf("loading big font: %s", err)
	}

	c := &Canvas{
		width:     width,
		height:    height,
		smallFont: smallFont,
		bigFont:   bigFont,
	}

	return c, nil
}

func (c *Canvas) Tick() image.Image {
	dc := gg.NewContext(c.width, c.height)
	dc.SetRGB(0, 0, 0)
	dc.DrawRectangle(0, 0, float64(c.width), float64(c.height))
	dc.Fill()

	switch c.CurrentScreen() {
	case CurrentTime:
		c.DrawCurrentTimeScreen(dc)
	case MinSecRemaining:
		c.DrawMinSecRemaining(dc)
	}

	return dc.Image()
}

func (c *Canvas) CurrentScreen() Screen {
	return MinSecRemaining
}
