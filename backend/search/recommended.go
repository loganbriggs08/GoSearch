package search

import (
	"github.com/NotKatsu/GoSearch/backend"
	"github.com/NotKatsu/GoSearch/database"
	"github.com/pterm/pterm"
)

func GetRecommended() []backend.FileReturnStruct {
	recommendedApps, recommendedAppsError := database.GetRecommendedApps()

	if recommendedAppsError != nil {
		pterm.Fatal.WithFatal(true).Println(recommendedAppsError)
	}

	return recommendedApps

}
