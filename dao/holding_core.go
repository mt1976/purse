package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/holding.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Holding (holding)
// Endpoint 	        : Holding (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:36
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

// Holding_GetList() returns a list of all Holding records
func Holding_GetList() (int, []dm.Holding, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Holding_SQLTable)
	count, holdingList, _, _ := holding_Fetch(tsql)
	
	return count, holdingList, nil
}



// Holding_GetByID() returns a single Holding record
func Holding_GetByID(id string) (int, dm.Holding, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Holding_SQLTable)
	tsql = tsql + " WHERE " + dm.Holding_SQLSearchID + "='" + id + "'"
	_, _, holdingItem, _ := holding_Fetch(tsql)

	return 1, holdingItem, nil
}



// Holding_DeleteByID() deletes a single Holding record
func Holding_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Holding_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Holding_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Holding_Store() saves/stores a Holding record to the database
func Holding_Store(r dm.Holding,req *http.Request) error {

	err := holding_Save(r,Audit_User(req))

	return err
}

// Holding_StoreSystem() saves/stores a Holding record to the database
func Holding_StoreSystem(r dm.Holding) error {
	
	err := holding_Save(r,Audit_Host())

	return err
}

// holding_Save() saves/stores a Holding record to the database
func holding_Save(r dm.Holding,usr string) error {

    var err error

// If there are fields below, create the methods in dao\Holding_Impl.go








































  r.Type  = holding_Type_OverrideStore (r)
  r.Portfolio  = holding_Portfolio_OverrideStore (r)

  r.Instrument  = holding_Instrument_OverrideStore (r)



	logs.Storing("Holding",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Holding_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.Holding_SYSId, r.SYSId)
ts = addData(ts, dm.Holding_Code, r.Code)
ts = addData(ts, dm.Holding_Name, r.Name)
ts = addData(ts, dm.Holding_Type, r.Type)
ts = addData(ts, dm.Holding_Portfolio, r.Portfolio)
ts = addData(ts, dm.Holding_Description, r.Description)
ts = addData(ts, dm.Holding_Amount, r.Amount)
ts = addData(ts, dm.Holding_Instrument, r.Instrument)
ts = addData(ts, dm.Holding_Price, r.Price)
ts = addData(ts, dm.Holding_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.Holding_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.Holding_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.Holding_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.Holding_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.Holding_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.Holding_Symbol, r.Symbol)






	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Holding_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Holding_Delete(r.Code)
	das.Execute(tsql)



	return err

}


// holding_Fetch read all employees
func holding_Fetch(tsql string) (int, []dm.Holding, dm.Holding, error) {

	var recItem dm.Holding
	var recList []dm.Holding

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Holding_SYSId, "0")
   recItem.Code  = get_String(rec, dm.Holding_Code, "")
   recItem.Name  = get_String(rec, dm.Holding_Name, "")
   recItem.Type  = get_String(rec, dm.Holding_Type, "")
   recItem.Portfolio  = get_String(rec, dm.Holding_Portfolio, "")
   recItem.Description  = get_String(rec, dm.Holding_Description, "")
   recItem.Amount  = get_String(rec, dm.Holding_Amount, "")
   recItem.Instrument  = get_String(rec, dm.Holding_Instrument, "")
   recItem.Price  = get_String(rec, dm.Holding_Price, "")
   recItem.SYSCreated  = get_String(rec, dm.Holding_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Holding_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Holding_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Holding_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Holding_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Holding_SYSUpdatedHost, "")
   recItem.Symbol  = get_String(rec, dm.Holding_Symbol, "")






// If there are fields below, create the methods in dao\Holding_Impl.go








































  recItem.Type  = holding_Type_OverrideFetch (recItem)
  recItem.Portfolio  = holding_Portfolio_OverrideFetch (recItem)

  recItem.Instrument  = holding_Instrument_OverrideFetch (recItem)

	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Holding_NewID(r dm.Holding) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

