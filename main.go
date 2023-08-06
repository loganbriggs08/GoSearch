package main

import (
	"embed"

	"github.com/pterm/pterm"

	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	GoSearchApp := GoSearch()

	err := wails.Run(&options.App{
		Title:         "GoSearch",
		Width:         650,
		Height:        350,
		Frameless:     true,
		DisableResize: true,
		AlwaysOnTop:   true,
		StartHidden:   true,
		Assets:        assets,

		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableFramelessWindowDecorations: true,
		},
		OnStartup: GoSearchApp.startup,
		Bind: []interface{}{
			GoSearchApp,
		},
	})

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
}
