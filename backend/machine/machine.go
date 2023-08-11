package machine

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/NotKatsu/GoSearch/database"

	"github.com/pterm/pterm"
)

func OpenExecutable(executablePath string) bool {
	application := exec.Command(executablePath)
	applicationOpenError := application.Start()

	if applicationOpenError != nil {
		pterm.Fatal.WithFatal(true).Println(applicationOpenError)
		return false
	}

	applicationOpenError = application.Wait()

	if applicationOpenError != nil {
		pterm.Fatal.WithFatal(true).Println(applicationOpenError)
		return false
	}

	return true
}

func CacheSystem() bool {
	directoriesToCache, err := getDirectoriesToCache()
	if err != nil {
		return false
	}

	for _, directory := range directoriesToCache {
		err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				fileName := info.Name()
				fileExtension := filepath.Ext(fileName)
				database.InsertIntoCache(path, fileName, fileExtension)
			}

			return nil
		})

		if err != nil {
			return false
		}
	}
	return true
}

func getDirectoriesToCache() ([]string, error) {
	var directories []string

	directories = append(directories,
		filepath.Join(os.Getenv("HOMEDRIVE")+os.Getenv("HOMEPATH"), "Desktop"),
		filepath.Join(os.Getenv("USERPROFILE"), "Documents"),
		filepath.Join(os.Getenv("USERPROFILE"), "Pictures"),
		filepath.Join(os.Getenv("USERPROFILE"), "Downloads"),
		filepath.Join(os.Getenv("USERPROFILE"), "Music"),
		filepath.Join(os.Getenv("USERPROFILE"), "Videos"))

	return directories, nil
}
