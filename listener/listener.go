package listener

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pterm/pterm"
	"os"
	"path/filepath"
	"runtime"
)

func getRootPath() (string, error) {
	if runtime.GOOS == "windows" {
		homeDrive := os.Getenv("HOMEDRIVE")
		homePath := os.Getenv("HOMEPATH")

		return homeDrive + homePath, nil
	} else {
		return "/", nil
	}
}

func CreateWatcher() {
	pterm.Info.WithShowLineNumber(true).Println("Watcher is currently being created...")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
	defer watcher.Close()

	pterm.Success.Println("Watcher has been successfully created.")

	userRootPath, getRootPathError := getRootPath()

	if getRootPathError != nil {
		pterm.Fatal.WithFatal(true).Println(getRootPathError)
	}

	filePathWalkError := filepath.Walk(userRootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			watcherAddError := watcher.Add(path)

			if watcherAddError != nil {
				return watcherAddError
			}

			fmt.Println(path)

		}
		return nil
	})

	if filePathWalkError != nil {
		pterm.Fatal.WithFatal(true).Println(filePathWalkError)
	}

	done := make(chan bool)
	go handleFileChange(watcher, done)

	<-done
}

func handleFileChange(watcher *fsnotify.Watcher, done chan bool) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				done <- true
				return
			}

			fmt.Println(getRootPath())

			switch {
			case event.Op&fsnotify.Create == fsnotify.Create:
				pterm.Println("File created:", event.Name)
			case event.Op&fsnotify.Write == fsnotify.Write:
				pterm.Println("File modified:", event.Name)
			case event.Op&fsnotify.Remove == fsnotify.Remove:
				pterm.Println("File deleted:", event.Name)
			}
		case err := <-watcher.Errors:
			pterm.Fatal.WithFatal(true).Println(err)
		}
	}
}
