package machine

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

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

	if runtime.GOOS == "windows" {
		homeDrive := os.Getenv("HOMEDRIVE")
		homePath := os.Getenv("HOMEPATH")
		userProfile := os.Getenv("USERPROFILE")

		desktop := filepath.Join(homeDrive+homePath, "Desktop")
		documents := filepath.Join(userProfile, "Documents")
		pictures := filepath.Join(userProfile, "Pictures")
		downloads := filepath.Join(userProfile, "Downloads")
		music := filepath.Join(userProfile, "Music")
		videos := filepath.Join(userProfile, "Videos")

		directories = append(directories, desktop, documents, pictures, downloads, music, videos)
	} else {
		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		pictures := filepath.Join(homeDirectory, "Pictures")
		downloads := filepath.Join(homeDirectory, "Downloads")
		music := filepath.Join(homeDirectory, "Music")
		videos := filepath.Join(homeDirectory, "Videos")

		directories = append(directories, homeDirectory, pictures, downloads, music, videos)
	}

	return directories, nil
}
