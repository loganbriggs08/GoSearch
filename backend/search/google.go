package search

import (
	"strings"

	"github.com/NotKatsu/GoSearch/backend"
)

func GetGoogle(query string) backend.FileReturnStruct {
	var googleBaseURL string = "https://www.google.com/search?q="

	GoogleReturnStruct := backend.FileReturnStruct{
		Name:     strings.ToLower(query),
		Location: googleBaseURL + query,
		
	}
}
