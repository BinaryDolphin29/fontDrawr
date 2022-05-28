//NotoSansJP-Regular.otf
package fontDrawer

import (
	"image"
	"image/png"
	"log"
	"os"
	"testing"
	"unicode/utf8"

	"golang.org/x/image/font/opentype"
)

func TestInit(t *testing.T) {
	_, err := NewDrawer(&Config{
		FontPath:    "",
		FaceOptions: &opentype.FaceOptions{},
		Width:       0,
		Height:      0,
	})

	if err == nil {
		t.Fatal("wrong argument value, but error is nil")
	}
}

const (
	width  = 96
	height = 64
)

var content = []byte("魑魅魍魎")

func TestAll(t *testing.T) {
	cfg := Config{
		Width:       width,
		Height:      height,
		FontPath:    "./font/NotoSansJP-Regular.otf",
		FaceOptions: nil,
	}

	drawer, err := NewDrawer(&cfg)
	if err != nil {
		log.Fatalln(err)
	}

	size := width / float64(utf8.RuneCountInString(string(content)))

	drawer.AppendContent(content)
	if err := drawer.ChangeFontSize(size); err != nil {
		t.Fatal(err)
	}
	drawer.ChangeFaceColor(image.White)
	drawer.SetPositionCenter()

	screen := drawer.Draw()

	testPng, err := os.Create("drawTest.png")
	if err != nil {
		t.Fatal(err)
	}

	if err := png.Encode(testPng, screen); err != nil {
		t.Fatal(err)
	}
}
