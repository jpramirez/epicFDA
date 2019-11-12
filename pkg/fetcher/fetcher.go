package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	models "github.com/jpramirez/epicFDA/pkg/models"
)

//FetcherOne struct to keep order on the fetchercommands
type FetcherOne struct {
	config models.Config
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

//NewFetcher creates new instance of the fetcher
func NewFetcher(config models.Config, BuildVersion string, BuidTime string) (FetcherOne, error) {
	var fetcherOne FetcherOne
	log.Println("Starting Go Quiet Place ")
	log.Println("Version : " + BuildVersion)
	log.Println("Build Time : " + BuidTime)
	fetcherOne.config = config
	return fetcherOne, nil
}

//FetchFDA will fetch https://api.fda.gov/download.json
func (F *FetcherOne) FetchFDA() (models.FDADownload, error) {
	var err error
	download := new(models.FDADownload)
	err = getJSON(F.config.APIURL, download)
	//For now we print the results
	return *download, err
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {
	if _, err := os.Stat(filepath); err == nil {
		fmt.Println("File Exists") // We dont download again
		return nil
	}
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

//DownloadAnimalAndVeterinary will parse the drug dataset model and download all files into config file related folder
func (F *FetcherOne) DownloadAnimalAndVeterinary(animal models.AnimalAndVeterinary) error {
	var err error
	fmt.Println("Animal and Veterinary Dataset")
	fmt.Println("Downloading partitions ")
	fmt.Println("Date :", animal.Event.ExportDate)

	for _, i := range animal.Event.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/AnimalAndVeterinary" + "/" + animal.Event.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}
	return err

}

//DownloadFood will parse the drug dataset model and download all files into config file related folder
func (F *FetcherOne) DownloadFood(food models.Food) error {
	var err error
	fmt.Println("Food Dataset")
	fmt.Println("Downloading partitions ")
	fmt.Println("Date :", food.Event.ExportDate)

	for _, i := range food.Event.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Food/Event/" + "/" + food.Event.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range food.Enforcement.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Food/Enforcement/" + "/" + food.Enforcement.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}
	return err

}

//DownloadDrug will parse the drug dataset model and download all files into config file related folder
func (F *FetcherOne) DownloadDrug(drug models.Drug) error {
	var err error
	fmt.Println("Drug Dataset")
	fmt.Println("Downloading partitions ")
	fmt.Println("Date :", drug.Event.ExportDate)

	for _, i := range drug.Event.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Drug/Event/" + "/" + drug.Event.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range drug.Enforcement.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Drug/Enforcement/" + "/" + drug.Enforcement.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range drug.Label.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Drug/Label/" + "/" + drug.Label.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range drug.NDC.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Drug/NDC/" + "/" + drug.NDC.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}
	return err

}

//DownloadDevice will parse the device dataset model and download all files into config file related folder
func (F *FetcherOne) DownloadDevice(device models.Device) error {
	var err error
	fmt.Println("device Dataset")
	fmt.Println("Downloading partitions ")
	fmt.Println("Date :", device.Event.ExportDate)

	for _, i := range device.P510K.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Device/510K/" + "/" + device.P510K.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range device.Classification.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Device/Classification/" + "/" + device.Classification.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range device.Event.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Device/Event/" + "/" + device.Event.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range device.Enforcement.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Device/Enforcement/" + "/" + device.Enforcement.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range device.Recall.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Device/Recall/" + "/" + device.Recall.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range device.RegistrationListing.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Device/RegistrationListing/" + "/" + device.RegistrationListing.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range device.PMA.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Device/PMA/" + "/" + device.PMA.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}

	for _, i := range device.UDI.Partitions {

		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder + "/Device/UDI/" + "/" + device.UDI.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println("DisplayName: ", i.DisplayName)
		fmt.Println("Size: ", i.SizeMB)
		fmt.Println("File: ", i.File)

		downloadFile(folderPath+path.Base(i.File), i.File)
	}
	return err
}
