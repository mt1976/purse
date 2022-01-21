package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/datasource.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataSource (datasource)
// Endpoint 	        : DataSource (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:36
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type DataSource struct {

SYSId        string
Code        string
Name        string
URI        string
Description        string
Notes        string
SYSCreated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedBy        string
SYSUpdatedHost        string
AuthKey        string

}

const (
	DataSource_Title       = "DataSource"
	DataSource_SQLTable    = "dataSource"
	DataSource_SQLSearchID = "Code"
	DataSource_QueryString = "Code"
	///
	/// Handler Defintions
	///
	DataSource_Template     = "DataSource"
	DataSource_TemplateList = "DataSource_List"
	DataSource_TemplateView = "DataSource_View"
	DataSource_TemplateEdit = "DataSource_Edit"
	DataSource_TemplateNew  = "DataSource_New"
	///
	/// Handler Monitor Paths
	///
	DataSource_Path       = "/API/DataSource/"
	DataSource_PathList   = "/DataSourceList/"
	DataSource_PathView   = "/DataSourceView/"
	DataSource_PathEdit   = "/DataSourceEdit/"
	DataSource_PathNew    = "/DataSourceNew/"
	DataSource_PathSave   = "/DataSourceSave/"
	DataSource_PathDelete = "/DataSourceDelete/"
	///
	/// SQL Field Definitions
	///
	DataSource_SYSId   = "_id" // SYSId is a Int
	DataSource_Code   = "Code" // Code is a String
	DataSource_Name   = "Name" // Name is a String
	DataSource_URI   = "URI" // URI is a String
	DataSource_Description   = "Description" // Description is a String
	DataSource_Notes   = "Notes" // Notes is a String
	DataSource_SYSCreated   = "_created" // SYSCreated is a String
	DataSource_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	DataSource_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	DataSource_SYSUpdated   = "_updated" // SYSUpdated is a String
	DataSource_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	DataSource_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	DataSource_AuthKey   = "AuthKey" // AuthKey is a String

	/// Definitions End
)
