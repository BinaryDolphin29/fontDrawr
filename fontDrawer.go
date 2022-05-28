package fontDrawer

import (
	"image"
	"io/ioutil"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

type Config struct {
	FontPath string
	// FaceOptions If nil, default options will be used.
	FaceOptions *opentype.FaceOptions
	Width       int
	Height      int
}

type Drawer struct {
	Drawer   *font.Drawer
	font     *sfnt.Font
	img      *image.RGBA
	content  []byte
	faceOpts *opentype.FaceOptions
}

// NewDrawer Create new *Drawer struct.
func NewDrawer(c *Config) (*Drawer, error) {
	readFont, err := ioutil.ReadFile(c.FontPath)
	if err != nil {
		return nil, err
	}

	otf, err := opentype.Parse(readFont)
	if err != nil {
		return nil, err
	}

	if c.FaceOptions == nil {
		c.FaceOptions = &opentype.FaceOptions{
			Size:    16,
			DPI:     72,
			Hinting: font.HintingNone,
		}
	}

	face, err := opentype.NewFace(otf, c.FaceOptions)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))

	drawer := &font.Drawer{
		Dst:  img,
		Src:  image.White,
		Face: face,
	}

	return &Drawer{
		font:     otf,
		img:      img,
		Drawer:   drawer,
		faceOpts: c.FaceOptions,
	}, nil
}

// Draw Drawing content on an image.
func (d *Drawer) Draw() *image.RGBA {
	d.Drawer.DrawBytes(d.content)
	return d.img
}

// Bounds Return the Drawer.BoundBytes.
func (d *Drawer) Bounds() (fixed.Rectangle26_6, fixed.Int26_6) {
	return d.Drawer.BoundBytes(d.content)
}

// Measure Return the Drawer.MeasureBytes.
func (d *Drawer) Measure() fixed.Int26_6 {
	return d.Drawer.MeasureBytes(d.content)
}

// SetContent Append to the content.
func (d *Drawer) AppendContent(b []byte) {
	d.content = append(d.content, b...)
}

// ChageFontOptions Change the font size.
func (d *Drawer) ChangeFontSize(size float64) error {
	d.faceOpts.Size = size
	newFace, err := opentype.NewFace(d.font, d.faceOpts)

	if err != nil {
		return err
	}

	d.Drawer.Face = newFace
	return nil
}

// ChangeFontOptions Change the font Hinting.
func (d *Drawer) ChangeFontHinting(hinting font.Hinting) error {
	d.faceOpts.Hinting = hinting
	newFace, err := opentype.NewFace(d.font, d.faceOpts)

	if err != nil {
		return err
	}

	d.Drawer.Face = newFace
	return nil
}

// ChageFaceColor Change the face color.
func (d *Drawer) ChangeFaceColor(uni *image.Uniform) {
	d.Drawer.Src = uni
}

// CenterX Return the computed center from the content.
func (d *Drawer) CenterX() fixed.Int26_6 {
	return (fixed.I(d.img.Bounds().Max.X) - d.Measure()) / 2
}

// CenterY Return the computed center from the content.
func (d *Drawer) CenterY() fixed.Int26_6 {
	b, _ := d.Bounds()
	max := b.Max
	min := b.Min

	return ((fixed.I(d.img.Bounds().Max.Y) - (max.Y - min.Y)) / 2) + (max.Y - min.Y)
}

// SetPosition Set the font start position.
func (d *Drawer) SetPosition(x, y fixed.Int26_6) {
	d.Drawer.Dot.X = x
	d.Drawer.Dot.Y = y
}

// SetPositionCenter Set the x and y positions to the center.
func (d *Drawer) SetPositionCenter() {
	d.SetPosition(d.CenterX(), d.CenterY())
}

// SetCenterXand The position X is center, and set the position Y.
func (d *Drawer) SetCenterXand(y fixed.Int26_6) {
	d.SetPosition(d.CenterX(), y)
}

// SetCenterYand The position Y is center, and set the position X.
func (d *Drawer) SetCenterYand(x fixed.Int26_6) {
	d.SetPosition(x, d.CenterY())
}

// ClearContent clear the content.
func (d *Drawer) ClearContent() {
	d.content = []byte{}
}

// ClearImg Clear Only the image.
func (d *Drawer) ClearImg() {
	maxW := d.img.Bounds().Max.X
	maxH := d.img.Bounds().Max.Y

	for pixY := 0; pixY < maxH; pixY++ {
		for pixX := 0; pixX < maxW; pixX++ {
			d.img.Set(pixX, pixY, image.Transparent)
		}
	}
}

// ClearAll Clear the content and image.
func (d *Drawer) ClearImgAndCtnt() {
	d.ClearContent()
	d.ClearImg()
}
