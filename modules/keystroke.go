package modules

import (
	"context"
	
	"github.com/pterm/pterm"
	"github.com/nsf/termbox-go"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var currentWindowStateOpen bool

func KeystrokeListener(ctx context.Context) {
	err := termbox.Init()

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
	defer termbox.Close()

	for {
		event := termbox.PollEvent()

		if event.Type == termbox.EventKey {
			if event.Key == termbox.KeyEsc {
				if currentWindowStateOpen == true {
					runtime.Hide(ctx)
					currentWindowStateOpen = false
				}
			} else if event.Key == termbox.KeyF5 {
				runtime.Show(ctx)
				currentWindowStateOpen = true
			}
		}
	}
}