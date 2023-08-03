package main

import (
	"context"
	"fmt"

	"github.com/NotKatsu/GoSearch/modules"
	"github.com/pterm/pterm"

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

func (a *App) HandleButtonClickEvent(application any) {
	runtime.Hide(a.ctx)

	applicationMap, successfulAssertion  := application.(map[string]interface{})

	if successfulAssertion == true {
		name := applicationMap["Name"].(string)
		location := applicationMap["Location"].(string)
		visits := applicationMap["Visits"].(float64)

		fmt.Println("Name:", name)
		fmt.Println("Location:", location)
		fmt.Println("Visits:", visits)
	} else {
		pterm.Fatal.WithFatal(true).Println("Something went wrong while trying to complete a Assertion.")
	}
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
