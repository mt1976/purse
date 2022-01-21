package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/activitytype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ActivityType (activitytype)
// Endpoint 	        : ActivityType (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:35
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type ActivityType struct {

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
	ActivityType_Title       = "ActivityType"
	ActivityType_SQLTable    = "activityType"
	ActivityType_SQLSearchID = "Code"
	ActivityType_QueryString = "Code"
	///
	/// Handler Defintions
	///
	ActivityType_Template     = "ActivityType"
	ActivityType_TemplateList = "ActivityType_List"
	ActivityType_TemplateView = "ActivityType_View"
	ActivityType_TemplateEdit = "ActivityType_Edit"
	ActivityType_TemplateNew  = "ActivityType_New"
	///
	/// Handler Monitor Paths
	///
	ActivityType_Path       = "/API/ActivityType/"
	ActivityType_PathList   = "/ActivityTypeList/"
	ActivityType_PathView   = "/ActivityTypeView/"
	ActivityType_PathEdit   = "/ActivityTypeEdit/"
	ActivityType_PathNew    = "/ActivityTypeNew/"
	ActivityType_PathSave   = "/ActivityTypeSave/"
	ActivityType_PathDelete = "/ActivityTypeDelete/"
	///
	/// SQL Field Definitions
	///
	ActivityType_SYSId   = "_id" // SYSId is a Int
	ActivityType_Code   = "Code" // Code is a String
	ActivityType_Name   = "Name" // Name is a String
	ActivityType_Factor   = "Factor" // Factor is a String
	ActivityType_SYSCreated   = "_created" // SYSCreated is a String
	ActivityType_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	ActivityType_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	ActivityType_SYSUpdated   = "_updated" // SYSUpdated is a String
	ActivityType_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	ActivityType_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
