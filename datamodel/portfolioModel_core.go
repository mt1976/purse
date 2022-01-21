package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/portfoliomodel.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : PortfolioModel (portfoliomodel)
// Endpoint 	        : PortfolioModel (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:37
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type PortfolioModel struct {

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
	PortfolioModel_Title       = "PortfolioModel"
	PortfolioModel_SQLTable    = "portfolioModel"
	PortfolioModel_SQLSearchID = "Code"
	PortfolioModel_QueryString = "Code"
	///
	/// Handler Defintions
	///
	PortfolioModel_Template     = "PortfolioModel"
	PortfolioModel_TemplateList = "PortfolioModel_List"
	PortfolioModel_TemplateView = "PortfolioModel_View"
	PortfolioModel_TemplateEdit = "PortfolioModel_Edit"
	PortfolioModel_TemplateNew  = "PortfolioModel_New"
	///
	/// Handler Monitor Paths
	///
	PortfolioModel_Path       = "/API/PortfolioModel/"
	PortfolioModel_PathList   = "/PortfolioModelList/"
	PortfolioModel_PathView   = "/PortfolioModelView/"
	PortfolioModel_PathEdit   = "/PortfolioModelEdit/"
	PortfolioModel_PathNew    = "/PortfolioModelNew/"
	PortfolioModel_PathSave   = "/PortfolioModelSave/"
	PortfolioModel_PathDelete = "/PortfolioModelDelete/"
	///
	/// SQL Field Definitions
	///
	PortfolioModel_SYSId   = "_id" // SYSId is a Int
	PortfolioModel_Code   = "Code" // Code is a String
	PortfolioModel_Name   = "Name" // Name is a String
	PortfolioModel_SYSCreated   = "_created" // SYSCreated is a String
	PortfolioModel_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	PortfolioModel_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	PortfolioModel_SYSUpdated   = "_updated" // SYSUpdated is a String
	PortfolioModel_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	PortfolioModel_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
