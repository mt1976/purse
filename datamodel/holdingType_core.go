package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/holdingtype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : HoldingType (holdingtype)
// Endpoint 	        : HoldingType (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:36
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type HoldingType struct {

SYSId        string
Code        string
Name        string
Factor        string
SYSCreated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedBy        string
SYSUpdatedHost        string

}

const (
	HoldingType_Title       = "HoldingType"
	HoldingType_SQLTable    = "holdingType"
	HoldingType_SQLSearchID = "Code"
	HoldingType_QueryString = "Code"
	///
	/// Handler Defintions
	///
	HoldingType_Template     = "HoldingType"
	HoldingType_TemplateList = "HoldingType_List"
	HoldingType_TemplateView = "HoldingType_View"
	HoldingType_TemplateEdit = "HoldingType_Edit"
	HoldingType_TemplateNew  = "HoldingType_New"
	///
	/// Handler Monitor Paths
	///
	HoldingType_Path       = "/API/HoldingType/"
	HoldingType_PathList   = "/HoldingTypeList/"
	HoldingType_PathView   = "/HoldingTypeView/"
	HoldingType_PathEdit   = "/HoldingTypeEdit/"
	HoldingType_PathNew    = "/HoldingTypeNew/"
	HoldingType_PathSave   = "/HoldingTypeSave/"
	HoldingType_PathDelete = "/HoldingTypeDelete/"
	///
	/// SQL Field Definitions
	///
	HoldingType_SYSId   = "_id" // SYSId is a Int
	HoldingType_Code   = "Code" // Code is a String
	HoldingType_Name   = "Name" // Name is a String
	HoldingType_Factor   = "Factor" // Factor is a String
	HoldingType_SYSCreated   = "_created" // SYSCreated is a String
	HoldingType_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	HoldingType_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	HoldingType_SYSUpdated   = "_updated" // SYSUpdated is a String
	HoldingType_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	HoldingType_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
