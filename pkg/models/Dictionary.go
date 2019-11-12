package models


//Device is a primary structure for FDA api, information about devices.
type Device struct {
	P510K	Dataset `json:"510k"`
	Classification	Dataset `json:"classification"`
	Enforcement	Dataset `json:"enforcement"`
	Recall Dataset `json:"recall"`
	RegistrationListing Dataset `json:"registrationlisting"`
	PMA	Dataset `json:"pma"`
	UDI	Dataset `json:"udi"`
	Event	Dataset `json:"event"`
}

//Food is the primary FDA api structure 
type Food struct {
	Enforcement	Dataset `json:"enforcement"`
	Event	Dataset `json:"event"`
}

//AnimalAndVeterinary is the primary FDA api structure 
type AnimalAndVeterinary struct {
	Event	Dataset `json:"event"`
}

//Other is the primary FDA api structure 
type Other struct {
	NSDE 	Dataset `json:"nsde"`
}


//Drug is the primary FDA api structure 
type Drug struct {
	Enforcement	Dataset `json:"enforcement"`
	NDC 	Dataset `json:"ndc"`
	Event	Dataset `json:"event"`
	Label Dataset `json:"label"`
}