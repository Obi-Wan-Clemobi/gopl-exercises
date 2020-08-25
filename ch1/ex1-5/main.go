// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
// Exercise 1.5: Change the Lissasjous program's color palette to green on black, for added authenticity. To create
// web color #RRGGBB, use color.RGBA{0xRR, 0xGG, 0xBB, 0xff}, where each pair of hexadecimal digits represents the
// intensity of the red, green, or blue component of the pixel

// Exercise 1.6: Modify the Lissajous program to produce images in multiple colors by adding more values to palette
// and then displaying them by changing the third argument of setColorIndex in some interesting way
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

//!-main
// Packages not needed by version in book.

//!+main
var green = color.RGBA{0x00, 0xff, 0x00, 0xff}
var blue = color.RGBA{0x00, 0x00, 0xff, 0xff}
var red = color.RGBA{0xff, 0xff, 0x00, 0x00}
var palette = []color.Color{color.Black, color.White, green, blue, red}

// var palette = []color.Color{color.White, color.Black}

const (
	// whiteIndex = 0 // first color in palette
	// blackIndex = 1 // next color in palette
	blackIndex = 0 // first color in palette
	whiteIndex = 1
	greenIndex = 2
	blueIndex  = 3
	redIndex   = 4
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 4     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			xPos := size + int(x*size+0.5)
			yPos := size + int(y*size+0.5)
			i := uint8(1 + rand.Float64()*4.0) // randomize color index
			img.SetColorIndex(xPos, yPos, i)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
