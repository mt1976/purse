package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/portfolio.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Portfolio (portfolio)
// Endpoint 	        : Portfolio (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:36
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Portfolio struct {

SYSId        string
Code        string
Name        string
Type        string
DefaultModel        string
Description        string
SYSCreated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedBy        string
SYSUpdatedHost        string
Type_Lookup        string
DefaultModel_Lookup        string



}

const (
	Portfolio_Title       = "Portfolio"
	Portfolio_SQLTable    = "portfolio"
	Portfolio_SQLSearchID = "Code"
	Portfolio_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Portfolio_Template     = "Portfolio"
	Portfolio_TemplateList = "Portfolio_List"
	Portfolio_TemplateView = "Portfolio_View"
	Portfolio_TemplateEdit = "Portfolio_Edit"
	Portfolio_TemplateNew  = "Portfolio_New"
	///
	/// Handler Monitor Paths
	///
	Portfolio_Path       = "/API/Portfolio/"
	Portfolio_PathList   = "/PortfolioList/"
	Portfolio_PathView   = "/PortfolioView/"
	Portfolio_PathEdit   = "/PortfolioEdit/"
	Portfolio_PathNew    = "/PortfolioNew/"
	Portfolio_PathSave   = "/PortfolioSave/"
	Portfolio_PathDelete = "/PortfolioDelete/"
	///
	/// SQL Field Definitions
	///
	Portfolio_SYSId   = "_id" // SYSId is a Int
	Portfolio_Code   = "Code" // Code is a String
	Portfolio_Name   = "Name" // Name is a String
	Portfolio_Type   = "Type" // Type is a String
	Portfolio_DefaultModel   = "DefaultModel" // DefaultModel is a String
	Portfolio_Description   = "Description" // Description is a String
	Portfolio_SYSCreated   = "_created" // SYSCreated is a String
	Portfolio_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Portfolio_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Portfolio_SYSUpdated   = "_updated" // SYSUpdated is a String
	Portfolio_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Portfolio_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	Portfolio_Type_Lookup   = "Type_Lookup" // Type_Lookup is a String
	Portfolio_DefaultModel_Lookup   = "DefaultModel_Lookup" // DefaultModel_Lookup is a String



	/// Definitions End
)
