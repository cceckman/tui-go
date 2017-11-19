package tui_test

import (
	"github.com/marcusolsson/tui-go"
	"github.com/marcusolsson/tui-go/tuitest"
	"image"
	"testing"
)

var drawTextEditTests = []struct {
	test  string
	size  image.Point
	setup func() *tui.TextEdit
	want  string
}{
	{
		test: "Simple",
		size: image.Point{15, 5},
		setup: func() *tui.TextEdit {
			e := tui.NewTextEdit()
			e.SetText("Lorem ipsum dolor sit amet")
			e.SetWordWrap(true)
			return e
		},
		want: `
Lorem ipsum    
dolor sit amet 
...............
...............
...............
`,
	},
}

func TestTextEdit_Draw(t *testing.T) {
	for _, tt := range drawTextEditTests {
		tt := tt
		t.Run(tt.test, func(t *testing.T) {
			var surface *tuitest.Surface
			if tt.size.X == 0 && tt.size.Y == 0 {
				surface = tuitest.NewSurface(10, 5)
			} else {
				surface = tuitest.NewSurface(tt.size.X, tt.size.Y)
			}

			painter := tui.NewPainter(surface, tui.NewTheme())
			painter.Repaint(tt.setup())

			if surface.String() != tt.want {
				t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), tt.want)
			}
		})
	}
}
