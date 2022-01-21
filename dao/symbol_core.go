package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/symbol.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Symbol (symbol)
// Endpoint 	        : Symbol (Code)
// For Project          : github.com/mt1976/purse/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 20/01/2022 at 14:44:37
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

// Symbol_GetList() returns a list of all Symbol records
func Symbol_GetList() (int, []dm.Symbol, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Symbol_SQLTable)
	count, symbolList, _, _ := symbol_Fetch(tsql)
	
	return count, symbolList, nil
}



// Symbol_GetByID() returns a single Symbol record
func Symbol_GetByID(id string) (int, dm.Symbol, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Symbol_SQLTable)
	tsql = tsql + " WHERE " + dm.Symbol_SQLSearchID + "='" + id + "'"
	_, _, symbolItem, _ := symbol_Fetch(tsql)

	return 1, symbolItem, nil
}



// Symbol_DeleteByID() deletes a single Symbol record
func Symbol_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Symbol_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Symbol_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Symbol_Store() saves/stores a Symbol record to the database
func Symbol_Store(r dm.Symbol,req *http.Request) error {

	err := symbol_Save(r,Audit_User(req))

	return err
}

// Symbol_StoreSystem() saves/stores a Symbol record to the database
func Symbol_StoreSystem(r dm.Symbol) error {
	
	err := symbol_Save(r,Audit_Host())

	return err
}

// symbol_Save() saves/stores a Symbol record to the database
func symbol_Save(r dm.Symbol,usr string) error {

    var err error

// If there are fields below, create the methods in dao\Symbol_Impl.go

































	logs.Storing("Symbol",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Symbol_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.Symbol_SYSId, r.SYSId)
ts = addData(ts, dm.Symbol_Code, r.Code)
ts = addData(ts, dm.Symbol_Name, r.Name)
ts = addData(ts, dm.Symbol_Type, r.Type)
ts = addData(ts, dm.Symbol_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.Symbol_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.Symbol_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.Symbol_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.Symbol_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.Symbol_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.Symbol_Identifier, r.Identifier)
ts = addData(ts, dm.Symbol_DataSource, r.DataSource)
ts = addData(ts, dm.Symbol_StaticValue, r.StaticValue)
ts = addData(ts, dm.Symbol_Factor, r.Factor)
ts = addData(ts, dm.Symbol_DPS, r.DPS)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Symbol_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	Symbol_Delete(r.Code)
	das.Execute(tsql)



	return err

}


// symbol_Fetch read all employees
func symbol_Fetch(tsql string) (int, []dm.Symbol, dm.Symbol, error) {

	var recItem dm.Symbol
	var recList []dm.Symbol

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.Symbol_SYSId, "0")
   recItem.Code  = get_String(rec, dm.Symbol_Code, "")
   recItem.Name  = get_String(rec, dm.Symbol_Name, "")
   recItem.Type  = get_String(rec, dm.Symbol_Type, "")
   recItem.SYSCreated  = get_String(rec, dm.Symbol_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.Symbol_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.Symbol_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.Symbol_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.Symbol_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.Symbol_SYSUpdatedHost, "")
   recItem.Identifier  = get_String(rec, dm.Symbol_Identifier, "")
   recItem.DataSource  = get_String(rec, dm.Symbol_DataSource, "")
   recItem.StaticValue  = get_String(rec, dm.Symbol_StaticValue, "")
   recItem.Factor  = get_String(rec, dm.Symbol_Factor, "")
   recItem.DPS  = get_String(rec, dm.Symbol_DPS, "")
// If there are fields below, create the methods in dao\Symbol_Impl.go































	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Symbol_NewID(r dm.Symbol) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

