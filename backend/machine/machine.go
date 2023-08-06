package machine

import (
	"fmt"
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
	fmt.Println(path)
	database.InsertIntoCache(path, filepath.Base(path), filepath.Ext(path))
	return nil
}

func CacheSystem() bool {
	roots := getDiskRoots()

	for _, root := range roots {
		filepathWalkError := filepath.Walk(root, addDataToDatabase)
		if filepathWalkError != nil {
			pterm.Fatal.WithFatal(true).Println(filepathWalkError)
			return false
		} else {
			return true
		}
	}

	return true
}
