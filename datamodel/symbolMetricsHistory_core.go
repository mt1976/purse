package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/symbolmetricshistory.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SymbolMetricsHistory (symbolmetricshistory)
// Endpoint 	        : SymbolMetricsHistory (TSKey)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:38
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type SymbolMetricsHistory struct {

SYSId        string
TSKey        string
Symbol        string
Date        string
Price        string
URI        string
SYSCreated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedBy        string
SYSUpdatedHost        string
Type        string
Bid        string
Offer        string
RawPrice        string
RawBid        string
RawOffer        string

}

const (
	SymbolMetricsHistory_Title       = "SymbolMetricsHistory"
	SymbolMetricsHistory_SQLTable    = "symbolMetricsHistory"
	SymbolMetricsHistory_SQLSearchID = "TSKey"
	SymbolMetricsHistory_QueryString = "TSKey"
	///
	/// Handler Defintions
	///
	SymbolMetricsHistory_Template     = "SymbolMetricsHistory"
	SymbolMetricsHistory_TemplateList = "SymbolMetricsHistory_List"
	SymbolMetricsHistory_TemplateView = "SymbolMetricsHistory_View"
	SymbolMetricsHistory_TemplateEdit = "SymbolMetricsHistory_Edit"
	SymbolMetricsHistory_TemplateNew  = "SymbolMetricsHistory_New"
	///
	/// Handler Monitor Paths
	///
	SymbolMetricsHistory_Path       = "/API/SymbolMetricsHistory/"
	SymbolMetricsHistory_PathList   = "/SymbolMetricsHistoryList/"
	SymbolMetricsHistory_PathView   = "/SymbolMetricsHistoryView/"
	SymbolMetricsHistory_PathEdit   = "/SymbolMetricsHistoryEdit/"
	SymbolMetricsHistory_PathNew    = "/SymbolMetricsHistoryNew/"
	SymbolMetricsHistory_PathSave   = "/SymbolMetricsHistorySave/"
	SymbolMetricsHistory_PathDelete = "/SymbolMetricsHistoryDelete/"
	///
	/// SQL Field Definitions
	///
	SymbolMetricsHistory_SYSId   = "_id" // SYSId is a Int
	SymbolMetricsHistory_TSKey   = "TSKey" // TSKey is a String
	SymbolMetricsHistory_Symbol   = "Symbol" // Symbol is a String
	SymbolMetricsHistory_Date   = "Date" // Date is a String
	SymbolMetricsHistory_Price   = "Price" // Price is a String
	SymbolMetricsHistory_URI   = "URI" // URI is a String
	SymbolMetricsHistory_SYSCreated   = "_created" // SYSCreated is a String
	SymbolMetricsHistory_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	SymbolMetricsHistory_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	SymbolMetricsHistory_SYSUpdated   = "_updated" // SYSUpdated is a String
	SymbolMetricsHistory_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	SymbolMetricsHistory_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	SymbolMetricsHistory_Type   = "Type" // Type is a String
	SymbolMetricsHistory_Bid   = "Bid" // Bid is a String
	SymbolMetricsHistory_Offer   = "Offer" // Offer is a String
	SymbolMetricsHistory_RawPrice   = "RawPrice" // RawPrice is a String
	SymbolMetricsHistory_RawBid   = "RawBid" // RawBid is a String
	SymbolMetricsHistory_RawOffer   = "RawOffer" // RawOffer is a String

	/// Definitions End
)
