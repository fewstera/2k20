package nyecanvas

import (
	"fmt"
	"io/ioutil"

	"github.com/zachomedia/go-bdf"
	"golang.org/x/image/font"
)

func loadFont(fontname string) (font.Face, error) {
	fontInput, err := ioutil.ReadFile(fmt.Sprintf("./fonts/%s.bdf", fontname))
	if err != nil {
		return nil, fmt.Errorf("reading font file: %s", err)
	}

	font, err := bdf.Parse(fontInput)
	if err != nil {
		return nil, fmt.Errorf("parsing font: %s", err)
	}

	return font.NewFace(), nil
}
