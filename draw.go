package main

import (
	"bytes"
	"slices"

	"github.com/nsf/termbox-go"
)

var frame_index = 0
var color_index = 0

func reverse(lines [][]byte) [][]byte {
	slices.Reverse(lines)
	return lines
}

func frameLines(frame []byte, orientation string) [][]byte {
	lines := bytes.Split(frame, []byte{'\n'})
	if orientation == "aussie" {
		lines = reverse(lines)
	}
	return lines
}

// box is the part of the frames that holds anything but spaces, the same
// for every frame so the parrot does not wobble as it dances (Asciiverse)
type box struct{ top, bottom, left, right int }

func bounds(animation Animation, orientation string) box {
	b := box{top: -1}
	for _, frame := range animation.Frames {
		for row, line := range frameLines(frame, orientation) {
			text := bytes.TrimRight(line, " \r")
			if len(bytes.TrimLeft(text, " ")) == 0 {
				continue
			}
			left := len(text) - len(bytes.TrimLeft(text, " "))
			if b.top < 0 {
				b = box{row, row, left, len(text)}
			}
			b.top = min(b.top, row)
			b.bottom = max(b.bottom, row)
			b.left = min(b.left, left)
			b.right = max(b.right, len(text))
		}
	}
	return b
}

var frame_box *box

func draw(animation Animation, orientation string) {
	_ = termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	lines := frameLines(animation.Frames[frame_index], orientation)

	// Centred, standing on the bottom row of the terminal; termbox leaves
	// out cells beyond the edges when the terminal is smaller (Asciiverse)
	if frame_box == nil {
		b := bounds(animation, orientation)
		frame_box = &b
	}
	width, height := termbox.Size()
	dx := (width-(frame_box.right-frame_box.left))/2 - frame_box.left
	dy := height - 1 - frame_box.bottom

	for x, line := range lines {
		for y, cell := range line {
			if cell == '\r' {
				continue
			}
			termbox.SetCell(y+dx, x+dy, rune(cell), colors[color_index], termbox.ColorDefault)
		}
	}

	_ = termbox.Flush()
	frame_index++
	color_index++
	if frame_index >= len(animation.Frames) {
		frame_index = 0
	}
	if color_index >= len(colors) {
		color_index = 0
	}
}
