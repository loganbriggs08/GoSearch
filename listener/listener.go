package listener

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/pterm/pterm"
)

func getRootPath() (string, error) {
	if runtime.GOOS == "windows" {
		homeDrive := os.Getenv("HOMEDRIVE")
		homePath := os.Getenv("HOMEPATH")

		desktop := filepath.Join(homeDrive+homePath, "Desktop")

		return desktop, nil
	} else {
		return "/", nil
	}
}

func handleFileChangeEvent(watcher *fsnotify.Watcher, done chan bool) {

}

func CreateWatcher() {
	watcher, newWatcherError := fsnotify.NewWatcher()

	if newWatcherError != nil {
		pterm.Fatal.WithFatal(true).Println(newWatcherError)
	}
	defer watcher.Close()

	rootPath, getRootPathError := getRootPath()
}
