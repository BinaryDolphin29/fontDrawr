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
	Size     float64
	Width    int
	Height   int
}

type Drawer struct {
	*Config
	Drawer *font.Drawer

	font    truetype.Font
	img     image.RGBA
	content []byte
}

// NewDrawer Create new *Drawer struct.
func NewDrawer(c *Config) (*Drawer, error) {
	readFont, _ := ioutil.ReadFile(c.FontPath)
	f, err := truetype.Parse(readFont)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(f, &truetype.Options{
		Size:    c.Size,
		Hinting: font.HintingVertical,
	})

	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))

	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: face,
	}

	return &Drawer{
		font:   *f,
		img:    *img,
		Drawer: drawer,
	}, nil
}

// Draw Drawing content on an image.
func (c *Drawer) Draw() *image.RGBA {
	c.Drawer.DrawBytes(c.content)
	return &c.img
}

// Bounds Return the Drawer.BoundBytes.
func (c *Drawer) Bounds() (fixed.Rectangle26_6, fixed.Int26_6) {
	return c.Drawer.BoundBytes(c.content)
}

// Measure Return the Drawer.MeasureBytes.
func (c *Drawer) Measure() fixed.Int26_6 {
	return c.Drawer.MeasureBytes(c.content)
}

// SetContent Append to the content.
func (c *Drawer) SetContent(str []byte) {
	c.content = append(c.content, str...)
}

// CenterX Return the computed center from the content.
func (c *Drawer) CenterX() fixed.Int26_6 {
	return (fixed.I(c.Config.Width) - c.Measure()) / fixed.I(2)
}

// ChageFontOptions Changing Size and Hinting of the font.
func (c *Drawer) ChageFontOptions(size float64, hinting *font.Hinting) {
	c.Drawer.Face = truetype.NewFace(&c.font, &truetype.Options{
		Size:    size,
		Hinting: *hinting,
	})
}

// ChageFaceColor Change the face color.
func (c *Drawer) ChageFaceColor(uni *image.Uniform) {
	c.Drawer.Src = uni
}

// SetPosition Set the font start position.
func (c *Drawer) SetPosition(x, y fixed.Int26_6) {
	c.Drawer.Dot.X = x
	c.Drawer.Dot.Y = y
}

// ClearContent clear the content.
func (c *Drawer) ClearContent() {
	c.content = []byte{}
}

// ClearImg Clear Only the image.
func (c *Drawer) ClearImg() {
	for pixX := 0; pixX < c.Config.Width; pixX++ {
		for pixY := 0; pixY < c.Config.Height; pixY++ {
			c.img.Set(pixX, pixY, image.Transparent)
		}
	}
}

// ClearAll Clear the content and image.
func (c *Drawer) ClearAll() {
	c.ClearContent()
	c.ClearImg()
}
