package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//FDADownload is the entire callback determined by https://api.fda.gov/download.json
type FDADownload struct {
	Meta    Meta    `json:"meta"`
	Results Results `json:"results"`
}

//Meta contains meta information of the download data set from https://api.fda.gov/download.json
type Meta struct {
	Disclaimer  string `json:"disclaimer"`
	Terms       string `json:"terms"`
	License     string `json:"license"`
	LastUpdated string `json:"last_updated"`
}

//Results holds the file listing dataset of https://api.fda.gov/download.json
type Results struct {
	Device              Device              `json:"device"`
	Food                Food                `json:"food"`
	Other               Other               `json:"other"`
	Drug                Drug                `json:"drug"`
	AnimalAndVeterinary AnimalAndVeterinary `json:"animalandveterinary"`
}

//SaveIndex will save the FDA json file index in the destination folder for future reference or use.
func (R *FDADownload) SaveIndex(Destination string) error {
	file, _ := json.MarshalIndent(R, "", " ")
	folderPath := Destination
	os.MkdirAll(folderPath, os.ModePerm)
	_ = ioutil.WriteFile(folderPath+"/"+"index.json", file, 0644)

	return nil
}
