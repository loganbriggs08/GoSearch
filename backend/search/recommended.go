package search

import (
	"github.com/NotKatsu/GoSearch/database"
	"github.com/NotKatsu/GoSearch/modules"
	"github.com/pterm/pterm"
)

func GetRecommended() []modules.FileReturnStruct {
	recommendedApps, recommendedAppsError := database.GetRecommendedApps()

	if recommendedAppsError != nil {
		pterm.Fatal.WithFatal(true).Println(recommendedAppsError)
	}

	return recommendedApps

}
