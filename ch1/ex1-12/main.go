// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
// Exercise 1.12: Modify the Lissajous server read parameter values from the URL.
// For example, you might arrange it so that a URL like http://localhost:8000/?cyles=20
// sets the number of cycles to 20 instead of the default 5. User strconv.Atoi function
// to convert the string parameter into an integer. You can see its doucmentation with
// go doc strconv.Atoi
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

type lissajousParams struct {
	cycles int
	size   int
}

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	cyclesStr := r.URL.Query()["cycles"]
	sizeStr := r.URL.Query()["size"]

	cycles, _ := strconv.Atoi(cyclesStr[0])
	size, _ := strconv.Atoi(sizeStr[0])

	lissajous(w, lissajousParams{cycles: cycles, size: size})

}

//!-handler

//!+main
var green = color.RGBA{0x00, 0xff, 0x00, 0xff}
var blue = color.RGBA{0x00, 0x00, 0xff, 0xff}
var red = color.RGBA{0xff, 0xff, 0x00, 0x00}
var palette = []color.Color{color.Black, color.White, green, blue, red}

const (
	blackIndex = 0 // first color in palette
	whiteIndex = 1
	greenIndex = 2
	blueIndex  = 3
	redIndex   = 4
)

func lissajous(out io.Writer, params lissajousParams) {
	const (
		res     = 0.001 // angular resolution
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	cycles := params.cycles // number of complete x oscillator revolutions
	size := params.size

	if cycles < 1 {
		cycles = 1
	}
	if size < 100 {
		size = 100
	}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*int(size)+1, 2*int(size)+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			xPos := size + int(x*float64(size)+0.5)
			yPos := size + int(y*float64(size)+0.5)
			i := uint8(1 + rand.Float64()*4.0) // randomize color index
			img.SetColorIndex(xPos, yPos, i)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
