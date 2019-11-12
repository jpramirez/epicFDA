package models



//FDADownload is the entire callback determined by https://api.fda.gov/download.json
type FDADownload struct {
	Meta	Meta `json:"meta"`
	Results Results `json:"results"`
}

//Meta contains meta information of the download data set from https://api.fda.gov/download.json
type Meta struct { 
	Disclaimer 	string `json:"disclaimer"`
	Terms	string `json:"terms"`
	License 	string `json:"license"`
	LastUpdate	string `json:"last_update"`
}

//Results holds the file listing dataset of https://api.fda.gov/download.json
type Results struct {
	Device	Device `json:"device"`
	Food	Food `json:"food"`
	Other	Other	`json:"other"`
	Drug	Drug `json:"drug"`
	AnimalAndVeterinary	AnimalAndVeterinary `json:"animalandveterinary"`
}