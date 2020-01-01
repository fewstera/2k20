package nyecanvas

import (
	"fmt"
	"math"
	"time"

	"github.com/fogleman/gg"
)

func (c *Canvas) DrawCurrentTimeScreen(dc *gg.Context) {
	t := time.Now()
	//isRed := t.Second()%10 > 5
	dc.SetFontFace(c.bigFont)
	animationDuration := float64(time.Duration(5) * time.Second / time.Nanosecond)
	fadeProgress := math.Abs(math.Sin(math.Pi * float64(t.UnixNano()) / animationDuration))

	red := 1 - fadeProgress
	blue := fadeProgress
	dc.SetRGB(red, 0.5, blue)

	dc.DrawString(t.Format("15:04"), 1, 11)

	dc.Fill()
}

func (c *Canvas) DrawMinSecRemaining(dc *gg.Context) {
	now := time.Now()
	ttm := c.midnight.Sub(now)

	mins := int(math.Floor(ttm.Minutes()))
	secs := int(math.Floor(ttm.Seconds())) % 60

	dc.SetFontFace(c.smallFont)
	dc.SetRGB(1, 1, 1)
	dc.DrawString(fmt.Sprintf("% 3dm", mins), 2, 7)
	dc.DrawString(fmt.Sprintf(" %2ds", secs), 2, 16)

	dc.Fill()
}

type colourDisplay struct {
	bg []float64
	fg []float64
}

func (c *Canvas) DrawSecRemaining(dc *gg.Context) {
	now := time.Now()
	ttm := int(c.midnight.Sub(now).Seconds())

	colours := []colourDisplay{
		colourDisplay{
			fg: []float64{1, 1, 1},
			bg: []float64{0.19215686274509805, 0.10588235294117647, 0.5725490196078431},
		},
		colourDisplay{
			fg: []float64{1, 1, 1},
			bg: []float64{0, 0.30196078431372547, 0.25098039215686274},
		},
		colourDisplay{
			fg: []float64{1, 1, 1},
			bg: []float64{0.8666666666666667, 0.17254901960784313, 0},
		},
	}
	currentColour := colours[int((time.Now().Unix()/5))%len(colours)]
	secsString := fmt.Sprintf("%d", ttm)

	dc.SetRGB(currentColour.bg[0], currentColour.bg[1], currentColour.bg[2])
	dc.DrawRectangle(0, 0, float64(c.width), float64(c.height))
	dc.Fill()

	dc.SetFontFace(c.massiveFont)
	dc.SetRGB(currentColour.fg[0], currentColour.fg[1], currentColour.fg[2])
	w, _ := dc.MeasureString(secsString)
	x := (float64(c.width) / 2) - (w / 2)

	dc.DrawString(secsString, x, 15)
	dc.Fill()
}
