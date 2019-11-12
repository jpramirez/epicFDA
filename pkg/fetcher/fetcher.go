package fetcher


import (
	"log"
    "io"
    "net/http"
	"os"
	 "fmt"
 "time"
 "encoding/json"
 "path"
 "strings"

models "github.com/jpramirez/EpicFDA/pkg/models"
)

//FetcherOne struct to keep order on the fetchercommands
type FetcherOne struct {
	config	models.Config
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
func (F * FetcherOne) FetchFDA () (models.Results, error) {
	var err error
	download  := new(models.FDADownload) // or &Foo{}
	err = getJSON(F.config.APIURL, download)
	//For now we print the results
    return download.Results,err
}


// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

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
func (F *FetcherOne) DownloadAnimalAndVeterinary(animal models.AnimalAndVeterinary)  error {
	var err error
	fmt.Println ("Animal and Veterinary Dataset")
	fmt.Println ("Downloading partitions ")
	fmt.Println("Date :", animal.Event.ExportDate)

	for _,i :=range animal.Event.Partitions  {


		displayName := strings.Replace(i.DisplayName, " ", "", -1)
		folderPath := F.config.DataSetFolder +"/AnimalAndVeterinary" + "/" + animal.Event.ExportDate + "/" + displayName + "/"
		os.MkdirAll(folderPath, os.ModePerm)

		fmt.Println ("DisplayName: ",i.DisplayName)
		fmt.Println ("Size: ",i.SizeMB)
		fmt.Println ("File: ",i.File)
		
		downloadFile(folderPath + path.Base(i.File), i.File)
	}
	return err

}