package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/holding.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Holding (holding)
// Endpoint 	        : Holding (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:36
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Holding struct {

SYSId        string
Code        string
Name        string
Type        string
Portfolio        string
Description        string
Amount        string
Instrument        string
Price        string
SYSCreated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedBy        string
SYSUpdatedHost        string
Symbol        string
Type_Lookup        string
Portfolio_Lookup        string


Instrument_Lookup        string


}

const (
	Holding_Title       = "Holding"
	Holding_SQLTable    = "holding"
	Holding_SQLSearchID = "Code"
	Holding_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Holding_Template     = "Holding"
	Holding_TemplateList = "Holding_List"
	Holding_TemplateView = "Holding_View"
	Holding_TemplateEdit = "Holding_Edit"
	Holding_TemplateNew  = "Holding_New"
	///
	/// Handler Monitor Paths
	///
	Holding_Path       = "/API/Holding/"
	Holding_PathList   = "/HoldingList/"
	Holding_PathView   = "/HoldingView/"
	Holding_PathEdit   = "/HoldingEdit/"
	Holding_PathNew    = "/HoldingNew/"
	Holding_PathSave   = "/HoldingSave/"
	Holding_PathDelete = "/HoldingDelete/"
	///
	/// SQL Field Definitions
	///
	Holding_SYSId   = "_id" // SYSId is a Int
	Holding_Code   = "Code" // Code is a String
	Holding_Name   = "Name" // Name is a String
	Holding_Type   = "Type" // Type is a String
	Holding_Portfolio   = "Portfolio" // Portfolio is a String
	Holding_Description   = "Description" // Description is a String
	Holding_Amount   = "Amount" // Amount is a String
	Holding_Instrument   = "Instrument" // Instrument is a String
	Holding_Price   = "Price" // Price is a String
	Holding_SYSCreated   = "_created" // SYSCreated is a String
	Holding_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Holding_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Holding_SYSUpdated   = "_updated" // SYSUpdated is a String
	Holding_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Holding_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	Holding_Symbol   = "Symbol" // Symbol is a String
	Holding_Type_Lookup   = "Type_Lookup" // Type_Lookup is a String
	Holding_Portfolio_Lookup   = "Portfolio_Lookup" // Portfolio_Lookup is a String


	Holding_Instrument_Lookup   = "Instrument_Lookup" // Instrument_Lookup is a String


	/// Definitions End
)
