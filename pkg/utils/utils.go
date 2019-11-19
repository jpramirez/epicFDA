package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	models "github.com/jpramirez/epicFDA/pkg/models"
)

//LoadConfiguration loads the req
func LoadConfiguration(file string) (models.Config, error) {
	var config models.Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
}

//DebugMessage will just be use to ease message control.
func DebugMessage(message string) {
	fmt.Printf(message)
}

func ListFiles(root string) ([]string, error) {

	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	return files, err

}
