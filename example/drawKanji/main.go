package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"unicode/utf8"

	"github.com/plasticbit/fontDrawer"
)

func main() {
	const (
		width  = 96
		height = 64
	)

	var content = []byte("魑魅魍魎")

	cfg := fontDrawer.Config{
		Width:       width,
		Height:      height,
		FontPath:    "../../font/NotoSansJP-Regular.otf",
		FaceOptions: nil,
	}

	drawer, err := fontDrawer.NewDrawer(&cfg)
	if err != nil {
		log.Fatalln(err)
	}

	size := width / float64(utf8.RuneCountInString(string(content)))

	drawer.AppendContent(content)
	if err := drawer.ChangeFontSize(size); err != nil {
		log.Fatalln(err)
	}
	drawer.ChangeFaceColor(image.White)
	drawer.SetPositionCenter()

	screen := drawer.Draw()

	testPng, err := os.Create("drawTest.png")
	if err != nil {
		log.Fatalln(err)
	}

	if err := png.Encode(testPng, screen); err != nil {
		log.Fatalln(err)
	}

}
