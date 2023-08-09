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

func CreateWatcher() {
	var err error

	pterm.Println("Watching for file changes...")

	watcher, newWatcherError := fsnotify.NewWatcher()
	if newWatcherError != nil {
		pterm.Fatal.WithFatal(true).Println(newWatcherError)
	}
	defer watcher.Close()

	rootPath, getRootPathError := getRootPath()
	if getRootPathError != nil {
		pterm.Fatal.WithFatal(true).Println(getRootPathError)
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
	go handleFileChange(watcher, done)

	// Wait for the program to finish (you can add other logic here if needed)
	<-done
}

func handleFileChange(watcher *fsnotify.Watcher, done chan bool) {
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				// File has been modified/created
				pterm.Println("File modified:", event.Name)
				// Update the database with the new/modified file information
				// database.UpdateFile(event.Name, ...)
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				// File has been deleted
				pterm.Println("File deleted:", event.Name)
				// Remove the file information from the database
				// database.DeleteFile(event.Name)
			}
		case err := <-watcher.Errors:
			// Handle errors
			pterm.Println("Error:", err.Error())
		}
	}
	done <- true
}
