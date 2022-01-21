package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/portfoliotype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : PortfolioType (portfoliotype)
// Endpoint 	        : PortfolioType (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:37
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type PortfolioType struct {

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
	PortfolioType_Title       = "PortfolioType"
	PortfolioType_SQLTable    = "portfolioType"
	PortfolioType_SQLSearchID = "Code"
	PortfolioType_QueryString = "Code"
	///
	/// Handler Defintions
	///
	PortfolioType_Template     = "PortfolioType"
	PortfolioType_TemplateList = "PortfolioType_List"
	PortfolioType_TemplateView = "PortfolioType_View"
	PortfolioType_TemplateEdit = "PortfolioType_Edit"
	PortfolioType_TemplateNew  = "PortfolioType_New"
	///
	/// Handler Monitor Paths
	///
	PortfolioType_Path       = "/API/PortfolioType/"
	PortfolioType_PathList   = "/PortfolioTypeList/"
	PortfolioType_PathView   = "/PortfolioTypeView/"
	PortfolioType_PathEdit   = "/PortfolioTypeEdit/"
	PortfolioType_PathNew    = "/PortfolioTypeNew/"
	PortfolioType_PathSave   = "/PortfolioTypeSave/"
	PortfolioType_PathDelete = "/PortfolioTypeDelete/"
	///
	/// SQL Field Definitions
	///
	PortfolioType_SYSId   = "_id" // SYSId is a Int
	PortfolioType_Code   = "Code" // Code is a String
	PortfolioType_Name   = "Name" // Name is a String
	PortfolioType_SYSCreated   = "_created" // SYSCreated is a String
	PortfolioType_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	PortfolioType_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	PortfolioType_SYSUpdated   = "_updated" // SYSUpdated is a String
	PortfolioType_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	PortfolioType_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
