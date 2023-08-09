package search

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/NotKatsu/GoSearch/backend"
	"github.com/pterm/pterm"
)

func GetGoogle(query string) []backend.GoogleReturnStruct {
	var searchQuery string = "https://www.google.com/search?q=" + url.QueryEscape(strings.ToLower(query))

	response, httpGetError := http.Get(searchQuery)

	if httpGetError != nil {
		pterm.Fatal.WithFatal(true).Println(httpGetError)
	}
	defer response.Body.Close()

	newDocument, newDocuementFromReaderError := goquery.NewDocumentFromReader(response.Body)

	if newDocuementFromReaderError != nil {
		pterm.Fatal.WithFatal(true).Println(newDocuementFromReaderError)
	}

	var searchResults []backend.GoogleReturnStruct

	newDocument.Find("div.rc h3").Each(func(index int, item *goquery.Selection) {
		link := item.Parent().Parent().Find("a")
		href, _ := link.Attr("href")

		if strings.Contains(item.Text(), " - ") {
			if item.Text() != "" && href != "" {
				NewGoogleReturnStruct := backend.GoogleReturnStruct{
					Title: item.Text(),
					Link:  href,
				}

				searchResults = append(searchResults, NewGoogleReturnStruct)

				if len(searchResults) >= 15 {
					return
				}

			}
		}
	})

	return searchResults
}
