package ch1

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"testing"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0, 255, 0, 255},
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 0, 255, 255}}

const (
	blackIndex = 0
	greenIndex = 1
	redIndex   = 2
	blueIndex  = 3
)

//产生利萨如曲线
func TestLissajous(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	f, err := os.OpenFile("Lissajous.gif", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if os.IsNotExist(err) {
		f, err = os.Create("Lissajous.gif")
		if err != nil {
			fmt.Printf("cannot create file %s", "Lissajous.gif")
		}
	}
	lissajous(f)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		c := uint8(rand.Intn(3) + 1)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5),
				size+int(y*size+0.5),
				c, )
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}