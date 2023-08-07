package json

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/NotKatsu/GoSearch/backend/appdata"
	"github.com/pterm/pterm"
)

type Settings struct {
	SystemCached bool `json:"system_cached"`
}

func getFileLocation() string {
	appDataPath, _ := appdata.GetAppDataFolder()

	appDataFolderPath := filepath.Join(appDataPath, "GoSearch")
	settingsFileLocation := appDataFolderPath + "\\settings.json"

	return settingsFileLocation
}

func readSettingsFromFile() (*Settings, error) {
	file, err := os.Open(getFileLocation())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var settings Settings
	err = json.NewDecoder(file).Decode(&settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}

func writeSettingsToFile(systemCached bool) error {
	settings := &Settings{
		SystemCached: systemCached,
	}

	data, err := json.MarshalIndent(settings, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(getFileLocation(), data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func SystemCached() bool {
	settings, err := readSettingsFromFile()
	if err != nil {
		return false
	}
	return settings.SystemCached
}

func UpdateCachedSetting(value bool) {
	err := writeSettingsToFile(value)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
}
