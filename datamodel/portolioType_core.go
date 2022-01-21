package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/portoliotype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : PortolioType (portoliotype)
// Endpoint 	        : PortolioType (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 18/01/2022 at 18:24:10
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type PortolioType struct {

SYSId        string
Code        string
Name        string
SYSCreated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedBy        string
SYSUpdatedHost        string

}

const (
	PortolioType_Title       = "PortfolioType"
	PortolioType_SQLTable    = "portfolioType"
	PortolioType_SQLSearchID = "Code"
	PortolioType_QueryString = "Code"
	///
	/// Handler Defintions
	///
	PortolioType_Template     = "PortolioType"
	PortolioType_TemplateList = "PortolioType_List"
	PortolioType_TemplateView = "PortolioType_View"
	PortolioType_TemplateEdit = "PortolioType_Edit"
	PortolioType_TemplateNew  = "PortolioType_New"
	///
	/// Handler Monitor Paths
	///
	PortolioType_Path       = "/API/PortolioType/"
	PortolioType_PathList   = "/PortolioTypeList/"
	PortolioType_PathView   = "/PortolioTypeView/"
	PortolioType_PathEdit   = "/PortolioTypeEdit/"
	PortolioType_PathNew    = "/PortolioTypeNew/"
	PortolioType_PathSave   = "/PortolioTypeSave/"
	PortolioType_PathDelete = "/PortolioTypeDelete/"
	///
	/// SQL Field Definitions
	///
	PortolioType_SYSId   = "_id" // SYSId is a Int
	PortolioType_Code   = "Code" // Code is a String
	PortolioType_Name   = "Name" // Name is a String
	PortolioType_SYSCreated   = "_created" // SYSCreated is a String
	PortolioType_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	PortolioType_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	PortolioType_SYSUpdated   = "_updated" // SYSUpdated is a String
	PortolioType_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	PortolioType_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
