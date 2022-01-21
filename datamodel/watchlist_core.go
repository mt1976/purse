package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/watchlist.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Watchlist (watchlist)
// Endpoint 	        : Watchlist (Id)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:39
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Watchlist struct {

SYSId        string
Id        string
Code        string
Name        string
Symbol        string
Notes        string
SYSCreated        string
SYSUpdated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdatedBy        string
SYSUpdatedHost        string

}

const (
	Watchlist_Title       = "Watchlist"
	Watchlist_SQLTable    = "watchlist"
	Watchlist_SQLSearchID = "id"
	Watchlist_QueryString = "Id"
	///
	/// Handler Defintions
	///
	Watchlist_Template     = "Watchlist"
	Watchlist_TemplateList = "Watchlist_List"
	Watchlist_TemplateView = "Watchlist_View"
	Watchlist_TemplateEdit = "Watchlist_Edit"
	Watchlist_TemplateNew  = "Watchlist_New"
	///
	/// Handler Monitor Paths
	///
	Watchlist_Path       = "/API/Watchlist/"
	Watchlist_PathList   = "/WatchlistList/"
	Watchlist_PathView   = "/WatchlistView/"
	Watchlist_PathEdit   = "/WatchlistEdit/"
	Watchlist_PathNew    = "/WatchlistNew/"
	Watchlist_PathSave   = "/WatchlistSave/"
	Watchlist_PathDelete = "/WatchlistDelete/"
	///
	/// SQL Field Definitions
	///
	Watchlist_SYSId   = "_id" // SYSId is a Int
	Watchlist_Id   = "Id" // Id is a String
	Watchlist_Code   = "Code" // Code is a String
	Watchlist_Name   = "Name" // Name is a String
	Watchlist_Symbol   = "Symbol" // Symbol is a String
	Watchlist_Notes   = "Notes" // Notes is a String
	Watchlist_SYSCreated   = "_created" // SYSCreated is a String
	Watchlist_SYSUpdated   = "_updated" // SYSUpdated is a String
	Watchlist_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Watchlist_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Watchlist_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Watchlist_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String

	/// Definitions End
)
