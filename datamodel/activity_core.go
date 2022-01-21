package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/activity.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Activity (activity)
// Endpoint 	        : Activity (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:35
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Activity struct {

SYSId        string
Code        string
Symbol        string
Date        string
Type        string
Portfolio        string
Amount        string
Price        string
Notes        string
Holding        string
ValueAMT        string
ValueCCY        string
SYSCreated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedBy        string
SYSUpdatedHost        string

}

const (
	Activity_Title       = "Activity"
	Activity_SQLTable    = "activity"
	Activity_SQLSearchID = "Code"
	Activity_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Activity_Template     = "Activity"
	Activity_TemplateList = "Activity_List"
	Activity_TemplateView = "Activity_View"
	Activity_TemplateEdit = "Activity_Edit"
	Activity_TemplateNew  = "Activity_New"
	///
	/// Handler Monitor Paths
	///
	Activity_Path       = "/API/Activity/"
	Activity_PathList   = "/ActivityList/"
	Activity_PathView   = "/ActivityView/"
	Activity_PathEdit   = "/ActivityEdit/"
	Activity_PathNew    = "/ActivityNew/"
	Activity_PathSave   = "/ActivitySave/"
	Activity_PathDelete = "/ActivityDelete/"
	///
	/// SQL Field Definitions
	///
	Activity_SYSId   = "_id" // SYSId is a Int
	Activity_Code   = "Code" // Code is a String
	Activity_Symbol   = "Symbol" // Symbol is a String
	Activity_Date   = "Date" // Date is a String
	Activity_Type   = "Type" // Type is a String
	Activity_Portfolio   = "Portfolio" // Portfolio is a String
	Activity_Amount   = "Amount" // Amount is a String
	Activity_Price   = "Price" // Price is a String
	Activity_Notes   = "Notes" // Notes is a String
	Activity_Holding   = "Holding" // Holding is a String
	Activity_ValueAMT   = "ValueAMT" // ValueAMT is a String
	Activity_ValueCCY   = "ValueCCY" // ValueCCY is a String
	Activity_SYSCreated   = "_created" // SYSCreated is a String
	Activity_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Activity_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Activity_SYSUpdated   = "_updated" // SYSUpdated is a String
	Activity_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Activity_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
