package search

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	
	"github.com/NotKatsu/GoSearch/modules"
	"github.com/pterm/pterm"
)

var (
	ProgramFilesLocation    = "C:\\Program Files"
	ProgramFilesX86Location = "C:\\Program Files (x86)"
	ProgramData             = "C:\\ProgramData"
)

func GetApplications(query string) []modules.FileReturnStruct {
	var appStructArray []modules.FileReturnStruct

	popularApplicationLocations := [3]string{ProgramFilesLocation, ProgramFilesX86Location, ProgramData}

	checkFiles := func(path string, fileInformation os.FileInfo, checkFilesError error) error {

		if checkFilesError != nil {
			pterm.Fatal.WithFatal(false).Println(checkFilesError)
		}

		if !fileInformation.IsDir() && filepath.Ext(fileInformation.Name()) == ".exe" && strings.Contains(fileInformation.Name(), query) {
			fmt.Printf("Checking %s", fileInformation.Name())
			currentAppStruct := modules.FileReturnStruct{
				Name:         string(unicode.ToUpper([]rune(fileInformation.Name())[0])) + fileInformation.Name()[1:],
				Location:     filepath.Dir(path) + "\\" + fileInformation.Name(),
				IconLocation: "",
			}


			fmt.Println(currentAppStruct.Name, currentAppStruct.Location, currentAppStruct.IconLocation)

			appStructArray = append(appStructArray, currentAppStruct)
		}

		return nil
	}

	for _, path := range popularApplicationLocations {
		if len(appStructArray) >= 1 {
			return appStructArray
		}

		fileWalkError := filepath.Walk(path, checkFiles)

		if fileWalkError != nil {
			pterm.Fatal.WithFatal(false).Println(fileWalkError)
		}
	}

	return appStructArray
}
