package os

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/pterm/pterm"
)

func GetAppDataFolder() (string, error) {
	currentUser, err := user.Current()

	if err != nil {
		return "", err
	}

	appDataPath := filepath.Join(currentUser.HomeDir, "AppData", "Roaming")

	return appDataPath, nil
}

func CreateAppDataFolder() (string, error) {
	appDataPath, err := GetAppDataFolder()

	if err != nil {
		return "", err
	}

	appDataFolderPath := filepath.Join(appDataPath, "GoSearch")

	if _, err := os.Stat(appDataFolderPath); os.IsNotExist(err) {
		err = os.MkdirAll(appDataFolderPath, 0755)

		if err != nil {
			return "", err
		}
	}

	return appDataFolderPath, nil
}

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

func CacheSystem() {
	
}
