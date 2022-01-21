package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/symbolmetricshistory.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : SymbolMetricsHistory (symbolmetricshistory)
// Endpoint 	        : SymbolMetricsHistory (TSKey)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:38
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"log"
	
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/purse/core"
	das  "github.com/mt1976/purse/das"
	
	
	dm   "github.com/mt1976/purse/datamodel"
	logs   "github.com/mt1976/purse/logs"
)

// SymbolMetricsHistory_GetList() returns a list of all SymbolMetricsHistory records
func SymbolMetricsHistory_GetList() (int, []dm.SymbolMetricsHistory, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.SymbolMetricsHistory_SQLTable)
	count, symbolmetricshistoryList, _, _ := symbolmetricshistory_Fetch(tsql)
	
	return count, symbolmetricshistoryList, nil
}



// SymbolMetricsHistory_GetByID() returns a single SymbolMetricsHistory record
func SymbolMetricsHistory_GetByID(id string) (int, dm.SymbolMetricsHistory, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.SymbolMetricsHistory_SQLTable)
	tsql = tsql + " WHERE " + dm.SymbolMetricsHistory_SQLSearchID + "='" + id + "'"
	_, _, symbolmetricshistoryItem, _ := symbolmetricshistory_Fetch(tsql)

	return 1, symbolmetricshistoryItem, nil
}



// SymbolMetricsHistory_DeleteByID() deletes a single SymbolMetricsHistory record
func SymbolMetricsHistory_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.SymbolMetricsHistory_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.SymbolMetricsHistory_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// SymbolMetricsHistory_Store() saves/stores a SymbolMetricsHistory record to the database
func SymbolMetricsHistory_Store(r dm.SymbolMetricsHistory,req *http.Request) error {

	err := symbolmetricshistory_Save(r,Audit_User(req))

	return err
}

// SymbolMetricsHistory_StoreSystem() saves/stores a SymbolMetricsHistory record to the database
func SymbolMetricsHistory_StoreSystem(r dm.SymbolMetricsHistory) error {
	
	err := symbolmetricshistory_Save(r,Audit_Host())

	return err
}

// symbolmetricshistory_Save() saves/stores a SymbolMetricsHistory record to the database
func symbolmetricshistory_Save(r dm.SymbolMetricsHistory,usr string) error {

    var err error

// If there are fields below, create the methods in dao\SymbolMetricsHistory_Impl.go







































	logs.Storing("SymbolMetricsHistory",fmt.Sprintf("%s", r))

	if len(r.TSKey) == 0 {
		r.TSKey = SymbolMetricsHistory_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.SymbolMetricsHistory_SYSId, r.SYSId)
ts = addData(ts, dm.SymbolMetricsHistory_TSKey, r.TSKey)
ts = addData(ts, dm.SymbolMetricsHistory_Symbol, r.Symbol)
ts = addData(ts, dm.SymbolMetricsHistory_Date, r.Date)
ts = addData(ts, dm.SymbolMetricsHistory_Price, r.Price)
ts = addData(ts, dm.SymbolMetricsHistory_URI, r.URI)
ts = addData(ts, dm.SymbolMetricsHistory_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.SymbolMetricsHistory_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.SymbolMetricsHistory_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.SymbolMetricsHistory_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.SymbolMetricsHistory_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.SymbolMetricsHistory_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.SymbolMetricsHistory_Type, r.Type)
ts = addData(ts, dm.SymbolMetricsHistory_Bid, r.Bid)
ts = addData(ts, dm.SymbolMetricsHistory_Offer, r.Offer)
ts = addData(ts, dm.SymbolMetricsHistory_RawPrice, r.RawPrice)
ts = addData(ts, dm.SymbolMetricsHistory_RawBid, r.RawBid)
ts = addData(ts, dm.SymbolMetricsHistory_RawOffer, r.RawOffer)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.SymbolMetricsHistory_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	SymbolMetricsHistory_Delete(r.TSKey)
	das.Execute(tsql)



	return err

}


// symbolmetricshistory_Fetch read all employees
func symbolmetricshistory_Fetch(tsql string) (int, []dm.SymbolMetricsHistory, dm.SymbolMetricsHistory, error) {

	var recItem dm.SymbolMetricsHistory
	var recList []dm.SymbolMetricsHistory

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.SymbolMetricsHistory_SYSId, "0")
   recItem.TSKey  = get_String(rec, dm.SymbolMetricsHistory_TSKey, "")
   recItem.Symbol  = get_String(rec, dm.SymbolMetricsHistory_Symbol, "")
   recItem.Date  = get_String(rec, dm.SymbolMetricsHistory_Date, "")
   recItem.Price  = get_String(rec, dm.SymbolMetricsHistory_Price, "")
   recItem.URI  = get_String(rec, dm.SymbolMetricsHistory_URI, "")
   recItem.SYSCreated  = get_String(rec, dm.SymbolMetricsHistory_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.SymbolMetricsHistory_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.SymbolMetricsHistory_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.SymbolMetricsHistory_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.SymbolMetricsHistory_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.SymbolMetricsHistory_SYSUpdatedHost, "")
   recItem.Type  = get_String(rec, dm.SymbolMetricsHistory_Type, "")
   recItem.Bid  = get_String(rec, dm.SymbolMetricsHistory_Bid, "")
   recItem.Offer  = get_String(rec, dm.SymbolMetricsHistory_Offer, "")
   recItem.RawPrice  = get_String(rec, dm.SymbolMetricsHistory_RawPrice, "")
   recItem.RawBid  = get_String(rec, dm.SymbolMetricsHistory_RawBid, "")
   recItem.RawOffer  = get_String(rec, dm.SymbolMetricsHistory_RawOffer, "")
// If there are fields below, create the methods in dao\SymbolMetricsHistory_Impl.go





































	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func SymbolMetricsHistory_NewID(r dm.SymbolMetricsHistory) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

