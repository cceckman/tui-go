package tui_test

import (
	"github.com/marcusolsson/tui-go"
	"github.com/marcusolsson/tui-go/tuitest"
	"image"
	"testing"
)

func TestMask_Full(t *testing.T) {
	sz := image.Pt(10, 10)
	surface := tuitest.NewSurface(sz.X, sz.Y)

	p := tui.NewPainter(surface, tui.NewTheme())
	p.WithMask(image.Rect(0, 0, sz.X, sz.Y), func(p *tui.Painter) {
		p.WithMask(image.Rect(0, 0, sz.Y, sz.Y), func(p *tui.Painter) {
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, '█')
				}
			}
		})
	})

	want := `
██████████
██████████
██████████
██████████
██████████
██████████
██████████
██████████
██████████
██████████
`
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestMask_Inset(t *testing.T) {
	sz := image.Pt(10, 10)
	surface := tuitest.NewSurface(sz.X, sz.Y)

	p := tui.NewPainter(surface, tui.NewTheme())
	p.WithMask(image.Rect(0, 0, sz.X, sz.Y), func(p *tui.Painter) {
		p.WithMask(image.Rect(1, 1, 9, 9), func(p *tui.Painter) {
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, '█')
				}
			}
		})
	})

	want := `
..........
.████████.
.████████.
.████████.
.████████.
.████████.
.████████.
.████████.
.████████.
..........
`
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestMask_FirstCell(t *testing.T) {
	sz := image.Pt(10, 10)
	surface := tuitest.NewSurface(sz.X, sz.Y)

	p := tui.NewPainter(surface, tui.NewTheme())
	p.WithMask(image.Rect(0, 0, sz.X, sz.Y), func(p *tui.Painter) {
		p.WithMask(image.Rect(0, 0, 1, 1), func(p *tui.Painter) {
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, '█')
				}
			}
		})
	})

	want := `
█.........
..........
..........
..........
..........
..........
..........
..........
..........
..........
`
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestMask_LastCell(t *testing.T) {
	sz := image.Pt(10, 10)
	surface := tuitest.NewSurface(sz.X, sz.Y)

	p := tui.NewPainter(surface, tui.NewTheme())
	p.WithMask(image.Rect(0, 0, sz.X, sz.Y), func(p *tui.Painter) {
		p.WithMask(image.Rect(9, 9, 10, 10), func(p *tui.Painter) {
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, '█')
				}
			}
		})
	})

	want := `
..........
..........
..........
..........
..........
..........
..........
..........
..........
.........█
`
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}

func TestMask_MaskWithinEmptyMaskIsHidden(t *testing.T) {
	sz := image.Pt(10, 10)
	surface := tuitest.NewSurface(sz.X, sz.Y)

	p := tui.NewPainter(surface, tui.NewTheme())
	p.WithMask(image.Rect(0, 0, 0, 0), func(p *tui.Painter) {
		p.WithMask(image.Rect(1, 1, 9, 9), func(p *tui.Painter) {
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, '█')
				}
			}
		})
	})

	want := `
..........
..........
..........
..........
..........
..........
..........
..........
..........
..........
`
	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}
