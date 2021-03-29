package fontDrawer

import (
	"image"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Config struct {
	FontPath string
	Size float64
	Width   int
	Height  int
}

type Drawer struct {
	content string
	Drawer font.Drawer
	Face   font.Face
	font   truetype.Font
	img    image.RGBA
}

func NewDrawer(c *Config) (*F, error) {
	readFont, _ := ioutil.ReadFile(fontPath)
	f, pErr := truetype.Parse(readFont)
	if pErr != nil {
		return nil, pErr
	}

	face := truetype.NewFace(f, &truetype.Options{
		Size:    size,
		Hinting: font.HintingVertical,
	})

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: face,
	}

	return &Config{
		Width:  width,
		Height: height,
		Drawer: *drawer,
		Face:   face,
		font:   *f,
		img:    *img,
	}, nil
}

func (c *Config) Draw() image.RGBA {
	c.Drawer.DrawBytes([]byte(c.content))
	return c.img
}

func (c *Config) Bounds() (fixed.Rectangle26_6, fixed.Int26_6) {
	return c.Drawer.BoundBytes([]byte(c.content))
}

func (c *Config) Measure() fixed.Int26_6 {
	return c.Drawer.MeasureBytes([]byte(c.content))
}

func (c *Config) SetContent(str string) {
	c.content += str
}

func (c *Config) CenterX() fixed.Int26_6 {
	return (fixed.I(c.Width) - c.Measure()) / 2
}

func (c *Config) SetFontSize(size float64) {
	c.Face = truetype.NewFace(&c.font, &truetype.Options{
		Size:    size,
		Hinting: font.HintingVertical,
	})
	
	c.Drawer.Face = c.Face
}

func (c *Config) SetPosition(x, y fixed.Int26_6) {
	c.Drawer.Dot.X = x
	c.Drawer.Dot.Y = y
}

func (c *Config) ClearContent() {
	c.content = ""
}

func (c *Config) Clear() {
	for pixX := 0; pixX < c.Width; pixX++ {
		for pixY := 0; pixY < c.Height; pixY++ {
			c.img.Set(pixX, pixY, image.Transparent)
		}
	}
}

func (c *Config) ClearAll() {
	c.ClearContent()
	c.Clear()
}
