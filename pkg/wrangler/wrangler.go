package wrangler

// This package is meant to read zipped information, and load DAtabases, we need the storage module for this to work.

import (
	"path/filepath"
	v1 "github.com/jpramirez/epicFDA/pkg/api/v1"
	models "github.com/jpramirez/epicFDA/pkg/models"
	utils "github.com/jpramirez/epicFDA/pkg/utils"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/gocql/gocql"
)

type WranglerObj struct {
	Session *gocql.Session
}

func (W *WranglerObj) ReadJsonFoodEventFromFile(fileName string) {


	extension := filepath.Ext(fileName)
	if extension == ".zip" {
		utils.Unzip(fileName, filepath.Dir(fileName))
		fileName = utils.FilenameWithoutExtension(fileName)
	}
	
	var result models.JsonFoodEventResults
	file, _ := ioutil.ReadFile(fileName)
	err := json.Unmarshal(file, &result)
	if err != nil {
		log.Println("Error?")
	}
	for _, event := range result.Results {
		W.SaveFoodEvent(event)
	}
}

func (W *WranglerObj) ReadJsonFoodEnforcementFromFile(fileName string) {

	extension := filepath.Ext(fileName)
	if extension == ".zip" {
		utils.Unzip(fileName, filepath.Dir(fileName))
		fileName = utils.FilenameWithoutExtension(fileName)
	}
	
	var result models.JsonFoodEnformentResults
	fmt.Println(fileName)
	file, _ := ioutil.ReadFile(fileName)
	err := json.Unmarshal(file, &result)
	if err != nil {
		fmt.Println("Error? ", err)
	}
	for _, enforcement := range result.Results {
		W.SaveFoodEnforcement(enforcement)
	}
}

//FindDataSet will return a full path of the dataset to load,
//Returns
func (W *WranglerObj) FindDataSet(datadir string, filename string, DataSetType string) []string {
	var results []string

	files, _ := utils.ListFiles(datadir + "/" + DataSetType)
	for _, file := range files {
		if filename == "all" {
			extension := filepath.Ext(file)
			if extension == ".zip" {
				results = append(results, file)
			}
		} else {
			if filename == file {
				results = append(results, file)
			}
		}

	}
	return results
}

/*
,center_classification_date DATE
,report_date                DATE
,postal_code                int
,termination_date           DATE
,recall_initiation_date     DATE

*/
func (W *WranglerObj) SaveFoodEnforcement(foodEvent v1.FoodEnforcement) error {
	var gocqlUuid gocql.UUID

	var err error

	gocqlUuid = gocql.TimeUUID()

	_reportDate, err := time.Parse("20060102", foodEvent.ReportDate)
	foodEvent.ReportDate = _reportDate.Format("2006-01-02")
	_reportDate, err = time.Parse("20060102", foodEvent.CenterClassificationDate)
	foodEvent.CenterClassificationDate = _reportDate.Format("2006-01-02")
	_reportDate, err = time.Parse("20060102", foodEvent.TerminationDate)
	foodEvent.TerminationDate = _reportDate.Format("2006-01-02")
	_reportDate, err = time.Parse("20060102", foodEvent.RecallInitiationDate)
	foodEvent.RecallInitiationDate = _reportDate.Format("2006-01-02")

	// write data to Cassandra
	err = W.Session.Query(`
	  INSERT INTO FoodEnforcement (
		FoodEnforcementID 
		,classification
		,center_classification_date
		,report_date
		,postal_code
		,termination_date
		,recall_initiation_date
		,recall_number
		,city 
		,more_code_info
		,event_id 
		,distribution_pattern
		,recalling_firm
		,voluntary_mandated
		,state
		,reason_for_recall
		,initial_firm_notification
		,status
		,product_type
		,country
		,product_description
		,code_info
		,address_1
		,address_2
		,product_quantity
	  ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?,?, ?, ?, ?, ?,?, ?, ?, ?, ?)`,
		gocqlUuid, foodEvent.Classification, foodEvent.CenterClassificationDate, foodEvent.ReportDate, foodEvent.PostalCode, foodEvent.TerminationDate,
		foodEvent.RecallInitiationDate, foodEvent.RecallNumber, foodEvent.City, foodEvent.MoreCodeInfo, foodEvent.EventId, foodEvent.DistributionPattern, foodEvent.RecallingFirm,
		foodEvent.VoluntaryMandated, foodEvent.State, foodEvent.ReasonForRecall, foodEvent.InitialFirmNotification, foodEvent.Status, foodEvent.ProductType, foodEvent.Country,
		foodEvent.ProductDescription, foodEvent.CodeInfo, foodEvent.Address_1, foodEvent.Address_2, foodEvent.ProductQuantity).Exec()

	if err != nil {
		fmt.Println("Something Happened ", err)
	}

	return err
}


//SaveFoodEvent will save the struct into the different tables.
func (W *WranglerObj) SaveFoodEvent(foodEvent v1.FoodEvent) error {
	var err error
	var gocqlUuid gocql.UUID

	/*
CREATE TABLE epicfda.FoodEvent (
    FoodEventID UUID Primary Key
    ,ReportNumber   int
    ,DateCreated    DATE
    ,DateStarted    DATE
    ,PRIMARY KEY (ReportNumber)
);
	*/
	fmt.Println(foodEvent)
	
	gocqlUuid = gocql.TimeUUID()

	_reportDate, err := time.Parse("20060102", foodEvent.DateCreated)
	foodEvent.DateCreated = _reportDate.Format("2006-01-02")
	_reportDate, err = time.Parse("20060102", foodEvent.DateStarted)
	foodEvent.DateStarted = _reportDate.Format("2006-01-02")

	err = W.Session.Query(`
	INSERT INTO epicfda.FoodEvent (
	FoodEventID,
	ReportNumber,
	DateCreated,
	DateStarted
	) VALUES (?, ?, ?,?)`,
	gocqlUuid,
	foodEvent.ReportNumber,
	foodEvent.DateCreated,
	foodEvent.DateStarted).Exec()

	if err != nil {
		fmt.Println("Something Happened inserting the event ", err)
	}

	// We inser the outcomes
	/*
	CREATE TABLE epicfda.FoodEventOutcomes (
		FoodEventID UUID Primary Key,
		Name    VARCHAR
	);
	*/

	for _, OutcomeName :=range (foodEvent.Outcomes) {
			err = W.Session.Query(`
			INSERT INTO epicfda.FoodEventOutcomes (
			FoodEventID,
			Name
			) VALUES (?, ?)`,
			gocqlUuid,
			OutcomeName).Exec()
			if err != nil {
				fmt.Println("Something Happened inserting outcomes ", err)
			}
	}


/*
CREATE TABLE epicfda.FoodEventReactions (
        FoodEventID  UUID Primary Key ,
        Name    VARCHAR
);
*/

	for _, reactions :=range (foodEvent.Reactions) {
		err = W.Session.Query(`
		INSERT INTO epicfda.FoodEventReactions (
		FoodEventID,
		Name
		) VALUES (?, ?)`,
		gocqlUuid,
		reactions).Exec()
		if err != nil {
			fmt.Println("Something Happened inserting reactions ", err)
		}
	}	
/*
CREATE TABLE epicfda.Product(
    FoodEventID   UUID Primary Key
    ,NameBrand             VARCHAR  
    ,IndustryCode           VARCHAR
    ,Role               VARCHAR
    ,IndustryName       VARCHAR
);

*/
	for _, product := range(foodEvent.Products) {
		err = W.Session.Query(`
		INSERT INTO epicfda.Product (
		FoodEventID,
		NameBrand,
		IndustryCode,
		Role,
		IndustryName
		) VALUES (?,?,?,?,?)`,
		gocqlUuid,
		product.NameBrand,
		product.IndustryCode,
		product.Role,
		product.IndustryName).Exec()

		if err != nil {
			fmt.Println("Something Happened inserting outcomes ", err)
		}
	}


/*
CREATE TABLE epicfda.Consumer(
    FoodEventID   UUID Primary Key
    ,Gender             VARCHAR  
    ,Age           VARCHAR
    ,AgeUnit               VARCHAR
);
*/
		err = W.Session.Query(`
		INSERT INTO epicfda.Consumer (
		FoodEventID,
		Gender,
		Age,
		AgeUnit
		) VALUES (?,?,?,?)`,
		gocqlUuid,
		foodEvent.Consumer.Gender,
		foodEvent.Consumer.Age,
		foodEvent.Consumer.AgeUnit).Exec()

		if err != nil {
			fmt.Println("Something Happened inserting Consumer ", err)
		}



	return err

}
