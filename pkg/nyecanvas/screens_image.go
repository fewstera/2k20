package nyecanvas

import (
	"github.com/fogleman/gg"
)

func (c *Canvas) DrawDisplayImage(dc *gg.Context) {
	dc.DrawImage(c.displayImage, 0, 0)
}
