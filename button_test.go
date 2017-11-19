package tui_test

import (
	"github.com/marcusolsson/tui-go"
	"github.com/marcusolsson/tui-go/tuitest"
	"testing"
)

func TestButton_OnActivated(t *testing.T) {
	btn := tui.NewButton("test")

	var invoked bool
	btn.OnActivated(func(b *tui.Button) {
		invoked = true
	})

	ev := tui.KeyEvent{
		Key: tui.KeyEnter,
	}

	t.Run("When button is not focused", func(t *testing.T) {
		btn.OnKeyEvent(ev)
		if invoked {
			t.Errorf("button should not be activated")
		}
	})

	invoked = false
	btn.SetFocused(true)

	t.Run("When button is focused", func(t *testing.T) {
		btn.OnKeyEvent(ev)
		if !invoked {
			t.Errorf("button should be activated")
		}
	})
}

func TestButton_Draw(t *testing.T) {
	surface := tuitest.NewSurface(10, 5)
	painter := tui.NewPainter(surface, tui.NewTheme())

	btn := tui.NewButton("test")
	painter.Repaint(btn)

	want := `
test      
..........
..........
..........
..........
`

	if surface.String() != want {
		t.Errorf("got = \n%s\n\nwant = \n%s", surface.String(), want)
	}
}
