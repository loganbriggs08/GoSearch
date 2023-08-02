package search

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/NotKatsu/GoSearch/database"
)

type RecommendedAppStruct struct {
	Name string
	Location string
	Visits uint16
}

func GetRecommended() []RecommendedAppStruct{
	recommendedApps, recommendedAppsError := database.GetRecommendedApps()
	
	if recommendedAppsError != nil {
		pterm.Fatal.WithFatal(true).Println(recommendedAppsError)
	}

	return recommendedApps

}