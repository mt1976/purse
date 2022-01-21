package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/symboltype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SymbolType (symboltype)
// Endpoint 	        : SymbolType (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:38
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type SymbolType struct {

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
	SymbolType_Title       = "SymbolType"
	SymbolType_SQLTable    = "symbolType"
	SymbolType_SQLSearchID = "Code"
	SymbolType_QueryString = "Code"
	///
	/// Handler Defintions
	///
	SymbolType_Template     = "SymbolType"
	SymbolType_TemplateList = "SymbolType_List"
	SymbolType_TemplateView = "SymbolType_View"
	SymbolType_TemplateEdit = "SymbolType_Edit"
	SymbolType_TemplateNew  = "SymbolType_New"
	///
	/// Handler Monitor Paths
	///
	SymbolType_Path       = "/API/SymbolType/"
	SymbolType_PathList   = "/SymbolTypeList/"
	SymbolType_PathView   = "/SymbolTypeView/"
	SymbolType_PathEdit   = "/SymbolTypeEdit/"
	SymbolType_PathNew    = "/SymbolTypeNew/"
	SymbolType_PathSave   = "/SymbolTypeSave/"
	SymbolType_PathDelete = "/SymbolTypeDelete/"
	///
	/// SQL Field Definitions
	///
	SymbolType_SYSId   = "_id" // SYSId is a Int
	SymbolType_Code   = "Code" // Code is a String
	SymbolType_Name   = "Name" // Name is a String
	SymbolType_SYSCreated   = "_created" // SYSCreated is a String
	SymbolType_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	SymbolType_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	SymbolType_SYSUpdated   = "_updated" // SYSUpdated is a String
	SymbolType_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	SymbolType_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
