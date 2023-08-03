package main

import (
	"context"
	"fmt"
	"github.com/NotKatsu/GoSearch/modules"

	"github.com/NotKatsu/GoSearch/database"
	"github.com/NotKatsu/GoSearch/modules/search"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/NotKatsu/GoSearch/modules/keystroke"
)

type App struct {
	ctx context.Context
}

func GoSearch() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	if database.SetupDatabase() == true {
		go keystroke.Listener(a.ctx)
	} else {
		runtime.Quit(a.ctx)
	}
}

func (a *App) HandleButtonClickEvent(result any) {
	fmt.Println(result)
}

func (a *App) Search(query string) []modules.RecommendedAppStruct{
	var	arrayWithEmptyStruct []modules.RecommendedAppStruct
	emptyStruct := modules.RecommendedAppStruct{}

	arrayWithEmptyStruct = append(arrayWithEmptyStruct, emptyStruct)

	if query == "" {
		return search.GetRecommended()
	} else if query != "" {
		return arrayWithEmptyStruct
	}

	return arrayWithEmptyStruct
}
