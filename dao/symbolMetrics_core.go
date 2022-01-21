package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/symbolmetrics.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : SymbolMetrics (symbolmetrics)
// Endpoint 	        : SymbolMetrics (Symbol)
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

// SymbolMetrics_GetList() returns a list of all SymbolMetrics records
func SymbolMetrics_GetList() (int, []dm.SymbolMetrics, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.SymbolMetrics_SQLTable)
	count, symbolmetricsList, _, _ := symbolmetrics_Fetch(tsql)
	
	return count, symbolmetricsList, nil
}



// SymbolMetrics_GetByID() returns a single SymbolMetrics record
func SymbolMetrics_GetByID(id string) (int, dm.SymbolMetrics, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.SymbolMetrics_SQLTable)
	tsql = tsql + " WHERE " + dm.SymbolMetrics_SQLSearchID + "='" + id + "'"
	_, _, symbolmetricsItem, _ := symbolmetrics_Fetch(tsql)

	return 1, symbolmetricsItem, nil
}



// SymbolMetrics_DeleteByID() deletes a single SymbolMetrics record
func SymbolMetrics_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.SymbolMetrics_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.SymbolMetrics_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// SymbolMetrics_Store() saves/stores a SymbolMetrics record to the database
func SymbolMetrics_Store(r dm.SymbolMetrics,req *http.Request) error {

	err := symbolmetrics_Save(r,Audit_User(req))

	return err
}

// SymbolMetrics_StoreSystem() saves/stores a SymbolMetrics record to the database
func SymbolMetrics_StoreSystem(r dm.SymbolMetrics) error {
	
	err := symbolmetrics_Save(r,Audit_Host())

	return err
}

// symbolmetrics_Save() saves/stores a SymbolMetrics record to the database
func symbolmetrics_Save(r dm.SymbolMetrics,usr string) error {

    var err error

// If there are fields below, create the methods in dao\SymbolMetrics_Impl.go

































	logs.Storing("SymbolMetrics",fmt.Sprintf("%s", r))

	if len(r.Symbol) == 0 {
		r.Symbol = SymbolMetrics_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.SymbolMetrics_SYSId, r.SYSId)
ts = addData(ts, dm.SymbolMetrics_Symbol, r.Symbol)
ts = addData(ts, dm.SymbolMetrics_Price, r.Price)
ts = addData(ts, dm.SymbolMetrics_URI, r.URI)
ts = addData(ts, dm.SymbolMetrics_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.SymbolMetrics_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.SymbolMetrics_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.SymbolMetrics_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.SymbolMetrics_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.SymbolMetrics_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.SymbolMetrics_Bid, r.Bid)
ts = addData(ts, dm.SymbolMetrics_Offer, r.Offer)
ts = addData(ts, dm.SymbolMetrics_RawPrice, r.RawPrice)
ts = addData(ts, dm.SymbolMetrics_RawBid, r.RawBid)
ts = addData(ts, dm.SymbolMetrics_RawOffer, r.RawOffer)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.SymbolMetrics_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	SymbolMetrics_Delete(r.Symbol)
	das.Execute(tsql)



	return err

}


// symbolmetrics_Fetch read all employees
func symbolmetrics_Fetch(tsql string) (int, []dm.SymbolMetrics, dm.SymbolMetrics, error) {

	var recItem dm.SymbolMetrics
	var recList []dm.SymbolMetrics

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.SymbolMetrics_SYSId, "0")
   recItem.Symbol  = get_String(rec, dm.SymbolMetrics_Symbol, "")
   recItem.Price  = get_String(rec, dm.SymbolMetrics_Price, "")
   recItem.URI  = get_String(rec, dm.SymbolMetrics_URI, "")
   recItem.SYSCreated  = get_String(rec, dm.SymbolMetrics_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.SymbolMetrics_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.SymbolMetrics_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.SymbolMetrics_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.SymbolMetrics_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.SymbolMetrics_SYSUpdatedHost, "")
   recItem.Bid  = get_String(rec, dm.SymbolMetrics_Bid, "")
   recItem.Offer  = get_String(rec, dm.SymbolMetrics_Offer, "")
   recItem.RawPrice  = get_String(rec, dm.SymbolMetrics_RawPrice, "")
   recItem.RawBid  = get_String(rec, dm.SymbolMetrics_RawBid, "")
   recItem.RawOffer  = get_String(rec, dm.SymbolMetrics_RawOffer, "")
// If there are fields below, create the methods in dao\SymbolMetrics_Impl.go































	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func SymbolMetrics_NewID(r dm.SymbolMetrics) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

