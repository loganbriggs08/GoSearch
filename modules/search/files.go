package search

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unicode"

	"github.com/NotKatsu/GoSearch/modules"
	"github.com/pterm/pterm"
)

var (
	ProgramFilesLocation    = "C:\\Program Files"
	ProgramFilesX86Location = "C:\\Program Files (x86)"
	ProgramData             = "C:\\ProgramData"
)

func Files(query string) []modules.FileReturnStruct {
	var appStructArray []modules.FileReturnStruct
	var wg sync.WaitGroup

	popularApplicationLocations := [3]string{ProgramFilesLocation, ProgramFilesX86Location, ProgramData}
	fileChan := make(chan modules.FileReturnStruct)

	checkFiles := func(path string) {
		defer wg.Done()

		filepath.Walk(path, func(path string, fileInformation os.FileInfo, err error) error {
			if err != nil {
				pterm.Fatal.WithFatal(false).Println(err)
			}

			if !fileInformation.IsDir() && filepath.Ext(fileInformation.Name()) == ".exe" && strings.Contains(fileInformation.Name(), query) {
				currentAppStruct := modules.FileReturnStruct{
					Name:     string(unicode.ToUpper([]rune(fileInformation.Name())[0])) + fileInformation.Name()[1:],
					Location: filepath.Join(filepath.Dir(path), fileInformation.Name()),
					}
					fileChan <- currentAppStruct
			}

			return nil
		})
	}

	for _, path := range popularApplicationLocations {
		wg.Add(1)
		go checkFiles(path)
	}

	go func() {
		wg.Wait()
		close(fileChan)
	}()

	for fileStruct := range fileChan {
		appStructArray = append(appStructArray, fileStruct)
	}

	return appStructArray
}