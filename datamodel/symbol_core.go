package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/symbol.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Symbol (symbol)
// Endpoint 	        : Symbol (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:37
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Symbol struct {

SYSId        string
Code        string
Name        string
Type        string
SYSCreated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedBy        string
SYSUpdatedHost        string
Identifier        string
DataSource        string
StaticValue        string
Factor        string
DPS        string

}

const (
	Symbol_Title       = "Symbol"
	Symbol_SQLTable    = "symbol"
	Symbol_SQLSearchID = "Code"
	Symbol_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Symbol_Template     = "Symbol"
	Symbol_TemplateList = "Symbol_List"
	Symbol_TemplateView = "Symbol_View"
	Symbol_TemplateEdit = "Symbol_Edit"
	Symbol_TemplateNew  = "Symbol_New"
	///
	/// Handler Monitor Paths
	///
	Symbol_Path       = "/API/Symbol/"
	Symbol_PathList   = "/SymbolList/"
	Symbol_PathView   = "/SymbolView/"
	Symbol_PathEdit   = "/SymbolEdit/"
	Symbol_PathNew    = "/SymbolNew/"
	Symbol_PathSave   = "/SymbolSave/"
	Symbol_PathDelete = "/SymbolDelete/"
	///
	/// SQL Field Definitions
	///
	Symbol_SYSId   = "_id" // SYSId is a Int
	Symbol_Code   = "Code" // Code is a String
	Symbol_Name   = "Name" // Name is a String
	Symbol_Type   = "Type" // Type is a String
	Symbol_SYSCreated   = "_created" // SYSCreated is a String
	Symbol_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Symbol_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Symbol_SYSUpdated   = "_updated" // SYSUpdated is a String
	Symbol_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Symbol_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	Symbol_Identifier   = "Identifier" // Identifier is a String
	Symbol_DataSource   = "DataSource" // DataSource is a String
	Symbol_StaticValue   = "StaticValue" // StaticValue is a String
	Symbol_Factor   = "Factor" // Factor is a String
	Symbol_DPS   = "DPS" // DPS is a String

	/// Definitions End
)
