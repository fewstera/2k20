package nyecanvas

import (
	"time"

	"github.com/fogleman/gg"
)

func (c *Canvas) DrawScrollingTextScreen(dc *gg.Context) {
	msg := c.textToDisplay
	dc.SetFontFace(c.bigFont)

	msgPixel := 32 + len(msg)*6

	runningTimeMS := int(time.Now().Sub(c.screenStartTime).Milliseconds())
	msPerPixel := 33

	scrollPosition := 32 - int(runningTimeMS/msPerPixel)%msgPixel

	dc.SetRGB(1, 0, 0)

	dc.DrawString(msg, float64(scrollPosition), 11)

	dc.Fill()
}

func (c *Canvas) StartDisplayingText(message string) {
	c.screenStartTime = time.Now()
	c.textDisplayEndTime = time.Now().Add(time.Duration(30) * time.Second)
	c.textToDisplay = message
}
