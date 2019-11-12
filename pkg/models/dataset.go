package models

//Dataset is the base structure for any result from the FDA page
type Dataset struct {
	DatasetName string `json:"dataset_name" db:"DatasetName"`
	TotalRecords string `json:"total_records" db:"TotalRecords"`
	ExportDate	string `json:"export_date" db:"ExportDate"`
	Partitions []partition `json:"partitions"`
}


