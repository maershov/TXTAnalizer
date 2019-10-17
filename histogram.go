package main

import (
	"flag"
	"fmt"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var col color.Color

const deltaX0 int = 100
const deltaY0 int = 100
const DELTA int = 100

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "./luxi-sans/Avanti Regular.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 12, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
)

func addText(img *image.RGBA, x, y int, text map[string]int) {
	flag.Parse()
	fontBytes, err := ioutil.ReadFile(*fontfile)
	var delta float64 = 50

	if err != nil {
		log.Println(err)
		return
	}

	f, err := freetype.ParseFont(fontBytes)

	if err != nil {
		log.Println(err)
		return
	}

	fg, bg := image.Black, image.White

	if *wonb {
		fg, bg = image.White, image.Black
	}

	rgba := img
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	switch *hinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	col = color.RGBA{0, 255, 0, 255} // Green
	Line(img, (0 + deltaX0), (10000 + deltaX0), (0 + deltaY0), (0 + deltaY0))
	VLine(img, (0 + deltaX0), (0 + deltaY0), (10000 + deltaY0))
	pt := freetype.Pt(100, 100)
	y0 := 100
	prevX := 100

	for s, v := range text {
		pt.Y += (c.PointToFixed(delta))
		_, err = c.DrawString(s+"       "+strconv.Itoa(v), pt)

		if err != nil {
			log.Println(err)
			return
		}

		col = color.RGBA{255, 0, 0, 255} // Red
		VLine(img, v*100+100, y0, y0+50)

		if prevX < 100+v*100 {
			HLine(img, y0, prevX, 100+v*100)
		} else {
			HLine(img, y0, 100+v*100, prevX)
		}

		prevX = 100 + v*100
		y0 += 50
	}
}

func VLine(img *image.RGBA, x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

func HLine(img *image.RGBA, y, x1, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

func Line(img *image.RGBA, x1, x2, y1, y2 int) {
	k := (y2 - y1) / (x2 - x1)
	b := y1 - k*x1
	for ; x1 <= x2; x1++ {
		y := k*x1 + b
		img.Set(x1, y, col)
	}
}

func Histogram(mapa map[string]int, lengthY, lengthX int) error {
	fmt.Println("size  ", lengthY, "    ", lengthX)
	//var img = image.NewRGBA(image.Rect(0, 0, lengthX*100 + 200,lengthY*50 + 200))   //картинка на компе долго обрабатывается при большом тексте
	var img = image.NewRGBA(image.Rect(0, 0, 10000, 10000))
	addText(img, 100, 100, mapa)

	f, err := os.Create("draw.png")

	if err != nil {
		return err
	}

	defer f.Close()

	png.Encode(f, img)

	return nil
}
