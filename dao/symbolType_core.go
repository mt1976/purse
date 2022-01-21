package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/symboltype.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : SymbolType (symboltype)
// Endpoint 	        : SymbolType (Code)
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

// SymbolType_GetList() returns a list of all SymbolType records
func SymbolType_GetList() (int, []dm.SymbolType, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.SymbolType_SQLTable)
	count, symboltypeList, _, _ := symboltype_Fetch(tsql)
	
	return count, symboltypeList, nil
}



// SymbolType_GetByID() returns a single SymbolType record
func SymbolType_GetByID(id string) (int, dm.SymbolType, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.SymbolType_SQLTable)
	tsql = tsql + " WHERE " + dm.SymbolType_SQLSearchID + "='" + id + "'"
	_, _, symboltypeItem, _ := symboltype_Fetch(tsql)

	return 1, symboltypeItem, nil
}



// SymbolType_DeleteByID() deletes a single SymbolType record
func SymbolType_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.SymbolType_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.SymbolType_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// SymbolType_Store() saves/stores a SymbolType record to the database
func SymbolType_Store(r dm.SymbolType,req *http.Request) error {

	err := symboltype_Save(r,Audit_User(req))

	return err
}

// SymbolType_StoreSystem() saves/stores a SymbolType record to the database
func SymbolType_StoreSystem(r dm.SymbolType) error {
	
	err := symboltype_Save(r,Audit_Host())

	return err
}

// symboltype_Save() saves/stores a SymbolType record to the database
func symboltype_Save(r dm.SymbolType,usr string) error {

    var err error

// If there are fields below, create the methods in dao\SymbolType_Impl.go





















	logs.Storing("SymbolType",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = SymbolType_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.SymbolType_SYSId, r.SYSId)
ts = addData(ts, dm.SymbolType_Code, r.Code)
ts = addData(ts, dm.SymbolType_Name, r.Name)
ts = addData(ts, dm.SymbolType_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.SymbolType_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.SymbolType_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.SymbolType_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.SymbolType_SYSUpdatedBy, r.SYSUpdatedBy)
ts = addData(ts, dm.SymbolType_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.SymbolType_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	SymbolType_Delete(r.Code)
	das.Execute(tsql)



	return err

}


// symboltype_Fetch read all employees
func symboltype_Fetch(tsql string) (int, []dm.SymbolType, dm.SymbolType, error) {

	var recItem dm.SymbolType
	var recList []dm.SymbolType

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 20/01/2022 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.SymbolType_SYSId, "0")
   recItem.Code  = get_String(rec, dm.SymbolType_Code, "")
   recItem.Name  = get_String(rec, dm.SymbolType_Name, "")
   recItem.SYSCreated  = get_String(rec, dm.SymbolType_SYSCreated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.SymbolType_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.SymbolType_SYSCreatedHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.SymbolType_SYSUpdated, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.SymbolType_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.SymbolType_SYSUpdatedHost, "")
// If there are fields below, create the methods in dao\SymbolType_Impl.go



















	// Automatically generated 20/01/2022 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func SymbolType_NewID(r dm.SymbolType) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

