package tui_test

import (
	"testing"

	"github.com/cceckman/tui-go"
	"github.com/cceckman/tui-go/tuitest"
	"github.com/kr/pretty"
)

func TestProgress_Draw(t *testing.T) {
	p := tui.NewProgress(100)
	p.SetSizePolicy(tui.Expanding, tui.Minimum)
	p.SetCurrent(50)

	surface := tuitest.NewSurface(11, 2)
	painter := tui.NewPainter(surface, tui.NewTheme())
	painter.Repaint(p)

	want := `
[===>-----]
...........
`

	if surface.String() != want {
		t.Error(pretty.Diff(surface.String(), want))
	}
}
