package main

import (
	"context"
	"fmt"

	"github.com/NotKatsu/GoSearch/backend"

	"github.com/NotKatsu/GoSearch/backend/dialog"
	"github.com/NotKatsu/GoSearch/backend/os"

	"github.com/pterm/pterm"

	"github.com/NotKatsu/GoSearch/backend/search"
	"github.com/NotKatsu/GoSearch/database"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/NotKatsu/GoSearch/backend/keystroke"
)

type App struct {
	ctx context.Context
}

func GoSearch() *App {
	return &App{}
}

var (
	currentPage = "Home"
)

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
	applicationMap, successfulAssertion := application.(map[string]interface{})

	if successfulAssertion == true {
		applicationName := applicationMap["Name"].(string)
		applicationLocation := applicationMap["Location"].(string)

		if os.OpenExecutable(applicationLocation) == false {
			errorMessage := "Failed to open " + applicationName
			dialog.ErrorDialog(errorMessage)
		}

	} else {
		pterm.Fatal.WithFatal(true).Println("Something went wrong while trying to complete a Assertion.")
	}
}

func (a *App) ToggleFavorite(name string, location string, favorite bool) []backend.FileReturnStruct {
	database.UpdateFavorite(name, location, favorite)

	return search.GetRecommended()
}

func (a *App) ClearCache() bool {
	if database.ClearDatabaseCache() == true {
		return true
	} else {
		return false
	}
}

func GetCurrentPage() string {
	return currentPage
}

func (a *App) Search(query string) []backend.FileReturnStruct {
	var arrayWithEmptyStruct []backend.FileReturnStruct
	emptyStruct := backend.FileReturnStruct{}

	arrayWithEmptyStruct = append(arrayWithEmptyStruct, emptyStruct)

	if query == "" {
		return search.GetRecommended()
	} else if query != "" {
		fmt.Println(query)
		//		return search.Files(strings.ToLower(query))
	}

	return arrayWithEmptyStruct
}
