package main

import (
	"fmt"
	"context"

	"github.com/NotKatsu/GoSearch/modules/dialog"
	"github.com/NotKatsu/GoSearch/modules/os"

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

func (a *App) Search(query string) []modules.FileReturnStruct {
	var arrayWithEmptyStruct []modules.FileReturnStruct
	emptyStruct := modules.FileReturnStruct{}

	arrayWithEmptyStruct = append(arrayWithEmptyStruct, emptyStruct)

	if query == "" {
		return search.GetRecommended()
	} else if query != "" {
		fmt.Println(query)
//		return search.Files(strings.ToLower(query))
	}

	return arrayWithEmptyStruct
}
