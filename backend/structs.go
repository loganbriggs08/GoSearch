package backend

type FileReturnStruct struct {
	Name         string
	Location     string
	IconLocation string
	Link         string
	Favorite     bool
	Visits       uint16
}

type GoogleReturnStruct struct {
	Title string
	Link  string
}
