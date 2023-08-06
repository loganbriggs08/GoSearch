package os

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/NotKatsu/GoSearch/database"

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

func getDiskRoots() []string {
	var roots []string
	if drives, err := filepath.Glob("/*"); err == nil {
		for _, drive := range drives {
			roots = append(roots, filepath.VolumeName(drive))
		}
	}
	return roots
}

func addDataToDatabase(path string, info os.FileInfo, err error) error {
	if err != nil {
		if os.IsPermission(err) {
			return nil
		} else {
			pterm.Fatal.WithFatal(true).Println(err)
			return err
		}
	}

	database.InsertIntoCache(path, filepath.Base(path), filepath.Ext(path))
	return nil
}

func CacheSystem() bool {
	roots := getDiskRoots()

	for _, root := range roots {
		filepathWalkError := filepath.Walk(root, addDataToDatabase)
		if filepathWalkError != nil {
			pterm.Fatal.WithFatal(true).Println(filepathWalkError)
		} else {
			return false
		}
	}
}
