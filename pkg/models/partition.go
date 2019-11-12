package models



type partition struct {
	SizeMB	string `json:"size_mb" db:"SizeMB"`
	Records string `json:"records" db:"Records"` // Amount of records in the dataset
	DisplayName	string `json:"display_name" db:"DisplayName"` 
	File	string `json:"file" db:"file"`
}