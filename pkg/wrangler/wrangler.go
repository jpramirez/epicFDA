package wrangler


// This package is meant to read zipped information, and load DAtabases, we need the storage module for this to work.

import (
	v1 "github.com/jpramirez/epicFDA/pkg/api/v1"
	models "github.com/jpramirez/epicFDA/pkg/models"

	"github.com/gocql/gocql"
	"encoding/json"
	"log"
	"time"
	"io/ioutil"
	"fmt"
)


type WranglerObj struct {
	Session *gocql.Session
}

func (W *WranglerObj) ReadJsonFoodEventFromFile (fileName string) {
	var result models.JsonFoodEnformentResults
	file, _ := ioutil.ReadFile(fileName)
	err := json.Unmarshal(file, &result)
	if err != nil {
		log.Println("Error?")
	}
	for _, enforcement := range result.Results{
		W.SaveFoodEnforcement(enforcement)
	}
}




func (W *WranglerObj) ReadJsonFoodEnforcementFromFile (fileName string) {
	var result models.JsonFoodEnformentResults
	file, _ := ioutil.ReadFile(fileName)
	err := json.Unmarshal(file, &result)
	if err != nil {
		log.Println("Error?")
	}
	for _, enforcement := range result.Results{
		W.SaveFoodEnforcement(enforcement)
	}
}

/*
,center_classification_date DATE  
,report_date                DATE  
,postal_code                int  
,termination_date           DATE  
,recall_initiation_date     DATE  
*/
func (W *WranglerObj) SaveFoodEnforcement (foodEvent v1.FoodEnforcement) error {
	var gocqlUuid gocql.UUID

	var err error

	gocqlUuid = gocql.TimeUUID()


	_reportDate,err:= time.Parse("20060102",foodEvent.ReportDate)
	foodEvent.ReportDate = _reportDate.Format("2006-01-02")
	_reportDate,err = time.Parse("20060102",foodEvent.CenterClassificationDate)
	foodEvent.CenterClassificationDate = _reportDate.Format("2006-01-02")
	_reportDate,err = time.Parse("20060102",foodEvent.TerminationDate)
	foodEvent.TerminationDate = _reportDate.Format("2006-01-02")
	_reportDate,err = time.Parse("20060102",foodEvent.RecallInitiationDate)
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
	  gocqlUuid, foodEvent.Classification, foodEvent.CenterClassificationDate,foodEvent.ReportDate, foodEvent.PostalCode, foodEvent.TerminationDate,
	  foodEvent.RecallInitiationDate,foodEvent.RecallNumber,foodEvent.City,foodEvent.MoreCodeInfo,foodEvent.EventId,foodEvent.DistributionPattern,foodEvent.RecallingFirm, 
	  foodEvent.VoluntaryMandated,foodEvent.State,foodEvent.ReasonForRecall,foodEvent.InitialFirmNotification,foodEvent.Status,foodEvent.ProductType,foodEvent.Country,
	  foodEvent.ProductDescription, foodEvent.CodeInfo,foodEvent.Address_1,foodEvent.Address_2,foodEvent.ProductQuantity).Exec()

	if (err!= nil){
		fmt.Println("Something Happened ", err)
	}


	return err
}


func (W *WranglerObj) SaveFoodEvent (foodEvent v1.FoodEvent) error {
	var err error

	return err
}