package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White}

const (
	twoPI  = math.Pi * 2
	colors = 255
)

func main() {
	setRainbow()
	lissajous(os.Stdout)
}

func setRainbow() {
	for i := 0; i < colors; i++ {
		red := math.Sin(float64(i)*twoPI/256.0+2.0)*128 + 127
		grn := math.Sin(float64(i)*twoPI/256.0+0.0)*128 + 127
		blu := math.Sin(float64(i)*twoPI/256.0+4.0)*128 + 127
		palette = append(palette, color.RGBA{uint8(red), uint8(grn), uint8(blu), 0xff})
	}
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
	colorIndex := 0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
			colorIndex++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
