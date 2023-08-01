package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"github.com/nsf/termbox-go"
	"github.com/pterm/pterm"
)

var currentWindowStateOpen bool

type App struct {
	ctx context.Context
}
func GoSearch() *App {
	return &App{}
}

func keystrokeListener(ctx context.Context) {
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

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	go keystrokeListener(a.ctx)
}


func (a *App) Search(name string) {
	fmt.Printf("Search Result: %s\n", name)
}
