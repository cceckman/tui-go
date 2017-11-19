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

func TestWithStyle_ApplyStyle(t *testing.T) {
	surface := tuitest.NewSurface(5, 5)

	theme := tui.NewTheme()
	theme.SetStyle("explicit", tui.Style{Fg: tui.ColorWhite, Bg: tui.ColorBlack})

	p := tui.NewPainter(surface, theme)
	p.WithMask(image.Rect(0, 0, 5, 5), func(p *tui.Painter) {
		p.WithMask(image.Rect(1, 1, 4, 4), func(p *tui.Painter) {
			sz := surface.Size()
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, ' ')
				}
			}

			p.WithMask(image.Rect(2, 2, 4, 4), func(p *tui.Painter) {
				p.WithStyle("explicit", func(p *tui.Painter) {
					sz := surface.Size()
					for x := 0; x < sz.X; x++ {
						for y := 0; y < sz.Y; y++ {
							p.DrawRune(x, y, '!')
						}
					}

				})
			})
		})
	})

	wantFg := `
.....
.000.
.022.
.022.
.....
`

	wantBg := `
.....
.000.
.011.
.011.
.....
`

	if surface.FgColors() != wantFg {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.FgColors(), wantFg)
	}
	if surface.BgColors() != wantBg {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.BgColors(), wantBg)
	}

}

func TestWithStyle_Stacks(t *testing.T) {
	surface := tuitest.NewSurface(10, 10)

	theme := tui.NewTheme()
	theme.SetStyle("explicit", tui.Style{Fg: tui.Color(3)})
	theme.SetStyle("auxiliary", tui.Style{Fg: tui.Color(2)})

	p := tui.NewPainter(surface, theme)
	p.WithMask(image.Rect(0, 0, 10, 10), func(p *tui.Painter) {

		// Set "explicit" and draw upper-left and upper-right.
		p.WithStyle("explicit", func(p *tui.Painter) {
			p.WithMask(image.Rect(1, 1, 4, 4), func(p *tui.Painter) {
				sz := surface.Size()
				for x := 0; x < sz.X; x++ {
					for y := 0; y < sz.Y; y++ {
						p.DrawRune(x, y, ' ')
					}
				}
			})
			// set "auxiliary" before drawing upper-right.
			p.WithStyle("auxiliary", func(p *tui.Painter) {
				p.WithMask(image.Rect(7, 1, 9, 3), func(p *tui.Painter) {
					sz := surface.Size()
					for x := 0; x < sz.X; x++ {
						for y := 0; y < sz.Y; y++ {
							p.DrawRune(x, y, ' ')
						}
					}
				})
			})
			// Then draw bottom-right, falling back to "explicit".
			p.WithMask(image.Rect(1, 6, 4, 9), func(p *tui.Painter) {
				sz := surface.Size()
				for x := 0; x < sz.X; x++ {
					for y := 0; y < sz.Y; y++ {
						p.DrawRune(x, y, ' ')
					}
				}
			})
		})

		// Use global default for bottom-right.
		p.WithMask(image.Rect(6, 6, 9, 9), func(p *tui.Painter) {
			sz := surface.Size()
			for x := 0; x < sz.X; x++ {
				for y := 0; y < sz.Y; y++ {
					p.DrawRune(x, y, ' ')
				}
			}
		})
	})

	wantFg := `
..........
.333...22.
.333...22.
.333......
..........
..........
.333..000.
.333..000.
.333..000.
..........
`

	if surface.FgColors() != wantFg {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.FgColors(), wantFg)
	}
}


