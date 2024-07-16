package main

import (
	"image"
	"image/color"
	"image/gif"
	"math/rand"
	"os"
)

const (
	width  = 1000
	height = 1000
	delay  = 1
)

func main() {

	// Create a new GIF file
	outFile, err := os.Create("out.gif")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	anim := gif.GIF{LoopCount: 0}

	data := initData(width, height)
	addFrame(&anim, data)

	for i := range height - 1 {
		addLine(data, i)
		addFrame(&anim, data)
	}

	err = gif.EncodeAll(outFile, &anim)
	if err != nil {
		panic(err)
	}

}

func addFrame(anim *gif.GIF, data [][]bool) {
	anim.Delay = append(anim.Delay, delay)

	frame := image.NewPaletted(image.Rect(0, 0, width, height), []color.Color{color.White, color.Black})
	for i := range data {
		for j := range data[i] {
			if data[i][j] {
				frame.SetColorIndex(j, i, 1) // set the pixel as Black
			}
		}
	}
	anim.Image = append(anim.Image, frame)
}

func initData(width, height int) [][]bool {

	data := make([][]bool, height)
	data[0] = make([]bool, width)

	for i := range data[0] {

		if random := rand.Float32(); random < 0.01 {
			data[0][i] = true
		}

	}
	return data
}

func addLine(data [][]bool, lastIndex int) {

	lastLine := data[lastIndex]
	data[lastIndex+1] = make([]bool, len(lastLine))

	for i := range lastLine {
		start := i - 1
		if start < 0 {
			start = 0
		}

		end := i + 1
		if end == len(lastLine) {
			end = len(lastLine) - 1
		}

		data[lastIndex+1][i] = checkRule(lastLine[start : end+1])
	}

}

func checkRule(vals []bool) bool {

	final := false
	attempts := 0
	for i := range vals {
		if vals[i] {
			attempts += 1
			final = true
			if attempts > 1 {
				final = false
				break
			}
		}

	}
	return final
}
