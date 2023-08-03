package modules

type RecommendedAppStruct struct {
	Name         string
	Location     string
	IconLocation string
	Favorite     bool
	Visits       uint16
}

type App struct {
	Name         string
	Location     string
	IconLocation string
}
