package models

//Config Main Configuration File Structure
type Config struct {
	WebPort      string   `json:"webport"`
	WebAddress   string   `json:"webaddress"`
	DatabaseName string   `json:"databasename"`
	APIURL       string   `json:"apiurl"`
	AppName      string   `json:"appname"`
	LogFile      string   `json:"logfile"`
	KeyFile      string   `json:"keyfile"`
	DataSetFolder string   `json:"datasetfolder"`
	CrtFile      string   `json:"crtfile"`
}
