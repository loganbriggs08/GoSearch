package main

import (
	"embed"

	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var assets embed.FS

func main() {
	GoSearchApp := GoSearch()

	err := wails.Run(&options.App{
		Title:  "GoSearch",
		Width:  900,
		Height: 350,
		Frameless: true,
		DisableResize: true,
		AlwaysOnTop: true,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent: true,
			DisableFramelessWindowDecorations: true,
		},
		OnStartup: GoSearchApp.startup,
		Bind: []interface{}{
			GoSearchApp,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
