package main

import (
	"context"
	"fmt"
	"strings"
	"time"

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
	currentPage = "Welcome"
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

func (a *App) ChangeTheme(theme string) {
	if database.SetTheme(theme) == true {
		currentPage = "Search"
		runtime.WindowReload(a.ctx)
	}
}

func (a *App) CurrentTheme() string {
	return database.GetCurrentTheme()
}

func (a *App) GetCurrentPage() string {
	return currentPage
}

func (a *App) SetPage(page string) {
	currentPage = page

	runtime.WindowReloadApp(a.ctx)
	runtime.Hide(a.ctx)
	time.Sleep(2 * time.Second)
	runtime.Show(a.ctx)
}

func (a *App) CloseApp() {
	runtime.Quit(a.ctx)
}

func (a *App) Search(query string) []backend.FileReturnStruct {
	var arrayWithEmptyStruct []backend.FileReturnStruct
	emptyStruct := backend.FileReturnStruct{}

	arrayWithEmptyStruct = append(arrayWithEmptyStruct, emptyStruct)

	if query == "" {
		return search.GetRecommended()
	} else if query != "" {
		if strings.ToLower(query) == "/settings" {
			currentPage = "Settings"
			runtime.WindowReload(a.ctx)
		}
		fmt.Println(query)
		//		return search.Files(strings.ToLower(query))
	}

	return arrayWithEmptyStruct
}
