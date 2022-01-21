package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/symbolmetrics.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SymbolMetrics (symbolmetrics)
// Endpoint 	        : SymbolMetrics (Symbol)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:38
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type SymbolMetrics struct {

SYSId        string
Symbol        string
Price        string
URI        string
SYSCreated        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdated        string
SYSUpdatedBy        string
SYSUpdatedHost        string
Bid        string
Offer        string
RawPrice        string
RawBid        string
RawOffer        string

}

const (
	SymbolMetrics_Title       = "SymbolMetrics"
	SymbolMetrics_SQLTable    = "symbolMetrics"
	SymbolMetrics_SQLSearchID = "Symbol"
	SymbolMetrics_QueryString = "Symbol"
	///
	/// Handler Defintions
	///
	SymbolMetrics_Template     = "SymbolMetrics"
	SymbolMetrics_TemplateList = "SymbolMetrics_List"
	SymbolMetrics_TemplateView = "SymbolMetrics_View"
	SymbolMetrics_TemplateEdit = "SymbolMetrics_Edit"
	SymbolMetrics_TemplateNew  = "SymbolMetrics_New"
	///
	/// Handler Monitor Paths
	///
	SymbolMetrics_Path       = "/API/SymbolMetrics/"
	SymbolMetrics_PathList   = "/SymbolMetricsList/"
	SymbolMetrics_PathView   = "/SymbolMetricsView/"
	SymbolMetrics_PathEdit   = "/SymbolMetricsEdit/"
	SymbolMetrics_PathNew    = "/SymbolMetricsNew/"
	SymbolMetrics_PathSave   = "/SymbolMetricsSave/"
	SymbolMetrics_PathDelete = "/SymbolMetricsDelete/"
	///
	/// SQL Field Definitions
	///
	SymbolMetrics_SYSId   = "_id" // SYSId is a Int
	SymbolMetrics_Symbol   = "Symbol" // Symbol is a String
	SymbolMetrics_Price   = "Price" // Price is a String
	SymbolMetrics_URI   = "URI" // URI is a String
	SymbolMetrics_SYSCreated   = "_created" // SYSCreated is a String
	SymbolMetrics_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	SymbolMetrics_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	SymbolMetrics_SYSUpdated   = "_updated" // SYSUpdated is a String
	SymbolMetrics_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	SymbolMetrics_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	SymbolMetrics_Bid   = "Bid" // Bid is a String
	SymbolMetrics_Offer   = "Offer" // Offer is a String
	SymbolMetrics_RawPrice   = "RawPrice" // RawPrice is a String
	SymbolMetrics_RawBid   = "RawBid" // RawBid is a String
	SymbolMetrics_RawOffer   = "RawOffer" // RawOffer is a String

	/// Definitions End
)
