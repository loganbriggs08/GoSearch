package keystroke

import (
	"context"

	"github.com/nsf/termbox-go"
	"github.com/pterm/pterm"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var CurrentWindowStateOpen bool

func Listener(ctx context.Context) {
	err := termbox.Init()

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
	defer termbox.Close()

	for {
		event := termbox.PollEvent()

		if event.Type == termbox.EventKey {
			if event.Key == termbox.KeyEsc {
				if CurrentWindowStateOpen == true {
					runtime.Hide(ctx)
					CurrentWindowStateOpen = false
				}
			} else if event.Key == termbox.KeyF5 {
				runtime.Show(ctx)
				CurrentWindowStateOpen = true
			}
		}
	}
}
