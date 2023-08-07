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
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				pterm.Println(event.Name)
			}

			if event.Op&fsnotify.Write == fsnotify.Remove {
				pterm.Println(event.Name)
			}
		}
	}
	done <- true
}

func CreateWatcher() {
	var err error

	watcher, newWatcherError := fsnotify.NewWatcher()

	if newWatcherError != nil {
		pterm.Fatal.WithFatal(true).Println(newWatcherError)
	}
	defer watcher.Close()

	rootPath, getRootPathError := getRootPath()

	if getRootPathError != nil {
		pterm.Fatal.WithFatal(true).Println(newWatcherError)
	}

	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
		return
	}

	done := make(chan bool)
	go handleFileChangeEvent(watcher, done)

	<-done
}
