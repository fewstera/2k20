package nyecanvas

import (
	"math"
	"time"

	"github.com/fogleman/gg"
)

func (c *Canvas) DrawCurrentTimeScreen(dc *gg.Context) Screen {
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
	return CurrentTime
}

func (c *Canvas) DrawMinSecRemaining(dc *gg.Context) Screen {
	t := time.Now()
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, 1)
	dc.SetFontFace(c.smallFont)

	dc.SetRGB(1, 1, 1)
	dc.DrawString(midnight.Format("15:04:05"), -1, 11)

	dc.Fill()
	return CurrentTime
}
