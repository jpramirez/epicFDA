package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"html/template"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	constants "github.com/jpramirez/epicFDA/pkg/constants"
	fetcher "github.com/jpramirez/epicFDA/pkg/fetcher"
	models "github.com/jpramirez/epicFDA/pkg/models"
	cassandra "github.com/jpramirez/epicFDA/pkg/storage"
	"github.com/jpramirez/epicFDA/pkg/wrangler"
)

type JResponse struct {
	ResponseCode string
	Message      string
	ResponseData []byte
}

type JResponseFileStatus struct {
	ResponseCode string
	Message      string
	FileStatus   []ResponseFileStatus
}

type ResponseFileStatus struct {
	FileName string
	Status   string
	Hash     string
}

//MainWebApp PHASE
type MainWebApp struct {
	Mux     *mux.Router
	Log     *log.Logger
	Config  models.Config
	Store   *sessions.CookieStore
	Storage *cassandra.StorageCassandra
}

//GetFileContentType will get the mime type of the file by reading its first 512 bytes (according to the standard)
func GetFileContentType(buffer []byte) (string, error) {
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

//NewApp creates a new instances
func NewApp(config models.Config) (MainWebApp, error) {

	var err error
	var wapp MainWebApp

	mux := mux.NewRouter().StrictSlash(true)
	f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log := log.New(os.Stdout, "web ", log.LstdFlags)

	wapp.Mux = mux
	wapp.Config = config
	wapp.Log = log
	wapp.Store = sessions.NewCookieStore([]byte("7b24afc8bc80e548d66c4e7ff72171c5"))
	var cs cassandra.StorageCassandra
	cs.Config = config
	cs.Init()
	wapp.Storage = &cs

	log.Println("NewAPP ---> Loggig Location")
	return wapp, err
}

func (M *MainWebApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	M.Mux.ServeHTTP(w, r)
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

//DownloadDeviceDataSet will download only the latest drug dataset
func (M *MainWebApp) DownloadDeviceDataSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	_fetch, err := fetcher.NewFetcher(M.Config, constants.BuildVersion, constants.BuildTime)
	if err != nil {
		log.Fatalln("Error on newebagent call ", err)
	}
	results, err := _fetch.FetchFDA()
	go _fetch.DownloadDevice(results.Results.Device)

	var response JResponse

	response.ResponseCode = "200 OK"
	response.Message = "Device Dataset Download Started"
	response.ResponseData = []byte("")
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)

}

//DownloadFoodDataSet will download only the latest drug dataset
func (M *MainWebApp) DownloadFoodDataSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	_fetch, err := fetcher.NewFetcher(M.Config, constants.BuildVersion, constants.BuildTime)
	if err != nil {
		log.Fatalln("Error on newebagent call ", err)
	}
	results, err := _fetch.FetchFDA()
	go _fetch.DownloadFood(results.Results.Food)

	var response JResponse

	response.ResponseCode = "200 OK"
	response.Message = "Food Dataset Download Started"
	response.ResponseData = []byte("")
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)

}

//DownloadAnimalDataSet will download only the latest drug dataset
func (M *MainWebApp) DownloadAnimalDataSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	_fetch, err := fetcher.NewFetcher(M.Config, constants.BuildVersion, constants.BuildTime)
	if err != nil {
		log.Fatalln("Error on newebagent call ", err)
	}
	results, err := _fetch.FetchFDA()
	go _fetch.DownloadAnimalAndVeterinary(results.Results.AnimalAndVeterinary)

	var response JResponse

	response.ResponseCode = "200 OK"
	response.Message = "Animal and Veterinary Dataset Download Started"
	response.ResponseData = []byte("")
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)

}

//DownloadDrugDataSet will download only the latest drug dataset
func (M *MainWebApp) DownloadDrugDataSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	_fetch, err := fetcher.NewFetcher(M.Config, constants.BuildVersion, constants.BuildTime)
	if err != nil {
		log.Fatalln("Error on newebagent call ", err)
	}
	results, err := _fetch.FetchFDA()
	go _fetch.DownloadDrug(results.Results.Drug)

	var response JResponse

	response.ResponseCode = "200 OK"
	response.Message = "Drug Dataset Download Started"
	response.ResponseData = []byte("")
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)

}

//DownloadIndex will download the entire json index from FDA
func (M *MainWebApp) DownloadIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	_fetch, err := fetcher.NewFetcher(M.Config, constants.BuildVersion, constants.BuildTime)
	if err != nil {
		log.Fatalln("Error on newebagent call ", err)
	}
	results, err := _fetch.FetchFDA()
	results.SaveIndex(M.Config.DataSetFolder)

	js, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//Liveness just keeps the connection alive
func (M *MainWebApp) Liveness(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var response JResponse

	response.ResponseCode = "200 OK"
	response.Message = "alive"
	response.ResponseData = []byte("")
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

func (M *MainWebApp) LoadFoodEnforcement(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	dataset := vars["dataset"]

	var wr wrangler.WranglerObj

	wr.Session = M.Storage.Session

	files := wr.FindDataSet(M.Config.DataSetFolder, dataset, "Food/Enforcement")

	/// HERE we need to replace with a search base on the file chosen by the url we can list for example  the zip files and decompress on demand.
	for _, f := range(files) {
		fmt.Println(f)
		wr.ReadJsonFoodEnforcementFromFile(f)
	}
	

	var response JResponse
	response.ResponseCode = "200 OK"
	response.Message = "alive"
	response.ResponseData = []byte("")
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}



func (M *MainWebApp) LoadFoodEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	dataset := vars["dataset"]

	var wr wrangler.WranglerObj

	wr.Session = M.Storage.Session

	files := wr.FindDataSet(M.Config.DataSetFolder, dataset, "Food/Event")

	/// HERE we need to replace with a search base on the file chosen by the url we can list for example  the zip files and decompress on demand.
	for _, f := range(files) {
		fmt.Println(f)
		wr.ReadJsonFoodEventFromFile(f)
	}
	

	var response JResponse
	response.ResponseCode = "200 OK"
	response.Message = "alive"
	response.ResponseData = []byte("")
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}

//HandleIndex for serving SPA
func (M *MainWebApp) HandleIndex(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("").ParseGlob("assets/templates/*.html"))
	t.ExecuteTemplate(w, "index.html", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
		"Stage":          os.Getenv("UP_STAGE"),
		"Year":           time.Now().Format("2006"),
		"EmojiCountry":   countryFlag(strings.Trim(r.Header.Get("Cloudfront-Viewer-Country"), "[]")),
	})
}

func getMetadata(r *http.Request) ([]byte, error) {
	f, _, err := r.FormFile("metadata")
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata form file: %v", err)
	}
	metadata, errRead := ioutil.ReadAll(f)
	if errRead != nil {
		return nil, fmt.Errorf("failed to read metadata: %v", errRead)
	}

	return metadata, nil
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	log.Println("setting up")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "admintoken, Content,Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func countryFlag(x string) string {
	if len(x) != 2 {
		return ""
	}
	if x[0] < 'A' || x[0] > 'Z' || x[1] < 'A' || x[1] > 'Z' {
		return ""
	}
	return string(0x1F1E6+rune(x[0])-'A') + string(0x1F1E6+rune(x[1])-'A')
}
