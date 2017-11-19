// Package tuitest provides utilities to help testing applications developed with gihtub.com/cceckman/tui-go.
package tuitest

import (
	"bytes"
	"image"

	"github.com/marcusolsson/tui-go"
)

type TestCell struct {
	Rune  rune
	Style tui.Style
}

// Surface implements the tui.Surface interface on top of a buffer, and
// provides additional methods to inspect its state after painting it.
type Surface struct {
	cells   map[image.Point]TestCell
	cursor  image.Point
	size    image.Point
	emptyCh rune
}

func NewSurface(w, h int) *Surface {
	return &Surface{
		cells:   make(map[image.Point]TestCell),
		size:    image.Point{w, h},
		emptyCh: '.',
	}
}

func (s *Surface) SetCell(x, y int, ch rune, style tui.Style) {
	s.cells[image.Point{x, y}] = TestCell{
		Rune:  ch,
		Style: style,
	}
}

func (s *Surface) SetCursor(x, y int) {
	s.cursor = image.Point{x, y}
}

func (s *Surface) HideCursor() {
	s.cursor = image.Point{}
}

func (s *Surface) Begin() {
	s.cells = make(map[image.Point]TestCell)
}

func (s *Surface) End() {
	// NOP
}

func (s *Surface) Size() image.Point {
	return s.size
}

func (s *Surface) String() string {
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
