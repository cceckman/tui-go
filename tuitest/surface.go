// Package tuitest provides utilities to help testing applications developed with gihtub.com/marcusolsson/tui-go.
package tuitest

import (
	"image"
	"github.com/marcusolsson/tui-go"
)

type TestCell struct {
	Rune  rune
	Style tui.Style
}

// TestSurface implements the tui.Surface interface on top of a buffer, and
// provides additional methods to inspect its state after painting it.
type TestSurface struct {
	cells   map[image.Point]TestCell
	cursor  image.Point
	size    image.Point
	emptyCh rune
}

func NewTestSurface(w, h int) *TestSurface {
	return &TestSurface{
		cells:   make(map[image.Point]TestCell),
		size:    image.Point{w, h},
		emptyCh: '.',
	}
}

func (s *TestSurface) SetCell(x, y int, ch rune, style tui.Style) {
	s.cells[image.Point{x, y}] = TestCell{
		Rune:  ch,
		Style: style,
	}
}

func (s *TestSurface) SetCursor(x, y int) {
	s.cursor = image.Point{x, y}
}

func (s *TestSurface) HideCursor() {
	s.cursor = image.Point{}
}

func (s *TestSurface) Begin() {
	s.cells = make(map[image.Point]TestCell)
}

func (s *TestSurface) End() {
	// NOP
}

func (s *TestSurface) Size() image.Point {
	return s.size
}

func (s *TestSurface) String() string {
	var buf bytes.Buffer
	buf.WriteRune('\n')
	for j := 0; j < s.size.Y; j++ {
		for i := 0; i < s.size.X; i++ {
			if cell, ok := s.cells[image.Point{i, j}]; ok {
				buf.WriteRune(cell.Rune)
			} else {
				buf.WriteRune(s.emptyCh)
			}
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}
